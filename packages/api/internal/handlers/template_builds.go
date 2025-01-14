package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/posthog/posthog-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/e2b-dev/infra/packages/api/internal/api"
	"github.com/e2b-dev/infra/packages/api/internal/constants"
	"github.com/e2b-dev/infra/packages/api/internal/nomad"
	"github.com/e2b-dev/infra/packages/api/internal/utils"
	"github.com/e2b-dev/infra/packages/shared/pkg/schema"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
)

func (a *APIStore) PostTemplates(c *gin.Context) {
	template := a.PostTemplatesWithoutResponse(c)
	if template != nil {
		c.JSON(http.StatusCreated, &template)
	}
}

func (a *APIStore) PostTemplatesWithoutResponse(c *gin.Context) *api.Template {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)

	// Prepare info for new env
	userID, team, tier, err := a.GetUserAndTeam(c)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, "Failed to get the default team")

		err = fmt.Errorf("error when getting default team: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	envID := utils.GenerateID()

	telemetry.SetAttributes(ctx,
		attribute.String("user.id", userID.String()),
		attribute.String("env.team.id", team.ID.String()),
		attribute.String("env.team.name", team.Name),
		attribute.String("env.id", envID),
		attribute.String("env.team.tier", tier.ID),
	)

	buildID, err := uuid.NewRandom()
	if err != nil {
		err = fmt.Errorf("error when generating build id: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		a.sendAPIStoreError(c, http.StatusInternalServerError, "Failed to generate build id")

		return nil
	}

	fileContent, fileHandler, err := c.Request.FormFile("buildContext")
	if err != nil {
		err = fmt.Errorf("error when parsing form data: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		a.sendAPIStoreError(c, http.StatusBadRequest, "Failed to parse form data")

		return nil
	}

	defer func() {
		closeErr := fileContent.Close()
		if closeErr != nil {
			errMsg := fmt.Errorf("error when closing file: %w", closeErr)

			telemetry.ReportError(ctx, errMsg)
		}
	}()

	// Check if file is a tar.gz file
	if !strings.HasSuffix(fileHandler.Filename, ".tar.gz.e2b") {
		err = fmt.Errorf("build context doesn't have correct extension, the file is %s", fileHandler.Filename)
		telemetry.ReportCriticalError(ctx, err)

		a.sendAPIStoreError(c, http.StatusBadRequest, "Build context must be a tar.gz.e2b file")

		return nil
	}

	dockerfile := c.PostForm("dockerfile")
	alias := c.PostForm("alias")
	startCmd := c.PostForm("startCmd")

	cpuCount, ramMB, apiError := getCPUAndRAM(tier.ID, c.PostForm("cpuCount"), c.PostForm("memoryMB"))
	if apiError != nil {
		telemetry.ReportCriticalError(ctx, apiError.Err)
		a.sendAPIStoreError(c, apiError.Code, apiError.ClientMsg)

		return nil
	}

	if alias != "" {
		alias, err = utils.CleanEnvID(alias)
		if err != nil {
			a.sendAPIStoreError(c, http.StatusBadRequest, fmt.Sprintf("Invalid alias: %s", alias))

			err = fmt.Errorf("invalid alias: %w", err)
			telemetry.ReportCriticalError(ctx, err)

			return nil
		}
	}

	properties := a.posthog.GetPackageToPosthogProperties(&c.Request.Header)
	a.posthog.IdentifyAnalyticsTeam(team.ID.String(), team.Name)
	a.posthog.CreateAnalyticsUserEvent(userID.String(), team.ID.String(), "submitted environment build request", properties.
		Set("environment", envID).
		Set("build_id", buildID).
		Set("dockerfile", dockerfile).
		Set("alias", alias),
	)

	telemetry.SetAttributes(ctx,
		attribute.String("build.id", buildID.String()),
		attribute.String("env.alias", alias),
		attribute.String("build.dockerfile", dockerfile),
		attribute.String("build.start_cmd", startCmd),
		attribute.Int64("build.cpuCount", cpuCount),
		attribute.Int64("build.ram_mb", ramMB),
	)

	_, err = a.cloudStorage.StreamFileUpload(strings.Join([]string{"v1", envID, buildID.String(), "context.tar.gz"}, "/"), fileContent)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when uploading file to cloud storage: %s", err))

		err = fmt.Errorf("error when uploading file to cloud storage: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	err = a.buildCache.Create(envID, buildID, team.ID)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusConflict, fmt.Sprintf("there's already running build for %s", envID))

		err = fmt.Errorf("build is already running build for %s", envID)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	telemetry.ReportEvent(ctx, "started creating new environment")

	if alias != "" {
		err = a.supabase.EnsureEnvAlias(ctx, alias, envID)
		if err != nil {
			a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when inserting alias: %s", err))

			err = fmt.Errorf("error when inserting alias: %w", err)
			telemetry.ReportCriticalError(ctx, err)

			a.buildCache.Delete(envID, buildID, team.ID)

			return nil
		} else {
			telemetry.ReportEvent(ctx, "inserted alias", attribute.String("env.alias", alias))
		}
	}

	go func() {
		buildContext, childSpan := a.tracer.Start(
			trace.ContextWithSpanContext(context.Background(), span.SpanContext()),
			"background-build-env",
		)

		var status api.TemplateBuildStatus

		buildErr := a.buildEnv(
			buildContext,
			userID.String(),
			team.ID,
			envID,
			buildID,
			dockerfile,
			startCmd,
			schema.DefaultKernelVersion,
			schema.DefaultFirecrackerVersion,
			properties,
			nomad.BuildConfig{
				VCpuCount:          cpuCount,
				MemoryMB:           ramMB,
				DiskSizeMB:         tier.DiskMB,
				KernelVersion:      schema.DefaultKernelVersion,
				FirecrackerVersion: schema.DefaultFirecrackerVersion,
			})

		if buildErr != nil {
			status = api.TemplateBuildStatusError

			errMsg := fmt.Errorf("error when building env: %w", buildErr)

			telemetry.ReportCriticalError(buildContext, errMsg)
		} else {
			status = api.TemplateBuildStatusReady

			telemetry.ReportEvent(buildContext, "created new environment", attribute.String("env.id", envID))
		}

		if status == api.TemplateBuildStatusError && alias != "" {
			errMsg := a.supabase.DeleteEnvAlias(buildContext, alias)
			if errMsg != nil {
				err = fmt.Errorf("error when deleting alias: %w", errMsg)
				telemetry.ReportError(buildContext, err)
			} else {
				telemetry.ReportEvent(buildContext, "deleted alias", attribute.String("env.alias", alias))
			}
		} else if status == api.TemplateBuildStatusReady && alias != "" {
			errMsg := a.supabase.UpdateEnvAlias(buildContext, alias, envID)
			if errMsg != nil {
				err = fmt.Errorf("error when updating alias: %w", errMsg)
				telemetry.ReportError(buildContext, err)
			} else {
				telemetry.ReportEvent(buildContext, "updated alias", attribute.String("env.alias", alias))
			}
		}

		cacheErr := a.buildCache.SetDone(envID, buildID, status)
		if cacheErr != nil {
			err = fmt.Errorf("error when setting build done in logs: %w", cacheErr)
			telemetry.ReportCriticalError(buildContext, cacheErr)
		}

		childSpan.End()
	}()

	aliases := []string{}

	if alias != "" {
		aliases = append(aliases, alias)
	}

	return &api.Template{
		TemplateID: envID,
		BuildID:    buildID.String(),
		Public:     false,
		Aliases:    &aliases,
	}
}

func (a *APIStore) PostTemplatesTemplateID(c *gin.Context, aliasOrTemplateID api.TemplateID) {
	template := a.PostTemplatesTemplateIDWithoutResponse(c, aliasOrTemplateID)

	if template != nil {
		c.JSON(http.StatusOK, &template)
	}
}

func (a *APIStore) PostTemplatesTemplateIDWithoutResponse(c *gin.Context, aliasOrTemplateID api.TemplateID) *api.Template {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)

	cleanedAliasOrEnvID, err := utils.CleanEnvID(aliasOrTemplateID)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusBadRequest, fmt.Sprintf("Invalid env ID: %s", aliasOrTemplateID))

		err = fmt.Errorf("invalid env ID: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	// Prepare info for rebuilding env
	userID, team, tier, err := a.GetUserAndTeam(c)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when getting default team: %s", err))

		err = fmt.Errorf("error when getting default team: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	telemetry.SetAttributes(ctx,
		attribute.String("user.id", userID.String()),
		attribute.String("env.team.id", team.ID.String()),
		attribute.String("env.team.tier", tier.ID),
	)

	buildID, err := uuid.NewRandom()
	if err != nil {
		err = fmt.Errorf("error when generating build id: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		a.sendAPIStoreError(c, http.StatusInternalServerError, "Failed to generate build id")

		return nil
	}

	fileContent, fileHandler, err := c.Request.FormFile("buildContext")
	if err != nil {
		err = fmt.Errorf("error when parsing form data: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		a.sendAPIStoreError(c, http.StatusBadRequest, "Failed to parse form data")

		return nil
	}

	// Check if file is a tar.gz file
	if !strings.HasSuffix(fileHandler.Filename, ".tar.gz.e2b") {
		err = fmt.Errorf("build context doesn't have correct extension, the file is %s", fileHandler.Filename)
		telemetry.ReportCriticalError(ctx, err)

		a.sendAPIStoreError(c, http.StatusBadRequest, "Build context must be a tar.gz.e2b file")

		return nil
	}

	dockerfile := c.PostForm("dockerfile")
	alias := c.PostForm("alias")

	cpuCount, ramMB, apiError := getCPUAndRAM(tier.ID, c.PostForm("cpuCount"), c.PostForm("memoryMB"))
	if apiError != nil {
		telemetry.ReportCriticalError(ctx, apiError.Err)
		a.sendAPIStoreError(c, apiError.Code, apiError.ClientMsg)

		return nil
	}

	if alias != "" {
		alias, err = utils.CleanEnvID(alias)
		if err != nil {
			a.sendAPIStoreError(c, http.StatusBadRequest, fmt.Sprintf("Invalid alias: %s", alias))

			err = fmt.Errorf("invalid alias: %w", err)
			telemetry.ReportCriticalError(ctx, err)

			return nil
		}
	}

	telemetry.SetAttributes(ctx,
		attribute.String("build.id", buildID.String()),
		attribute.Int64("build.cpuCount", cpuCount),
		attribute.Int64("build.ram_mb", ramMB),
		attribute.String("env.alias", alias),
	)

	env, envKernelVersion, envFirecrackerVersion, hasAccess, accessErr := a.CheckTeamAccessEnv(ctx, cleanedAliasOrEnvID, team.ID, false)
	if accessErr != nil {
		a.sendAPIStoreError(c, http.StatusNotFound, fmt.Sprintf("the sandbox template '%s' does not exist", cleanedAliasOrEnvID))

		errMsg := fmt.Errorf("env not found: %w", accessErr)
		telemetry.ReportError(ctx, errMsg)

		return nil
	}

	properties := a.posthog.GetPackageToPosthogProperties(&c.Request.Header)
	a.posthog.IdentifyAnalyticsTeam(team.ID.String(), team.Name)
	a.posthog.CreateAnalyticsUserEvent(userID.String(), team.ID.String(), "submitted environment build request", properties.
		Set("environment", env.TemplateID).
		Set("build_id", buildID).
		Set("alias", alias).
		Set("dockerfile", dockerfile),
	)

	if !hasAccess {
		a.sendAPIStoreError(c, http.StatusForbidden, "You don't have access to this sandbox template")

		errMsg := fmt.Errorf("user doesn't have access to env '%s'", env.TemplateID)
		telemetry.ReportError(ctx, errMsg)

		return nil
	}

	telemetry.SetAttributes(ctx, attribute.String("build.id", buildID.String()))

	_, err = a.cloudStorage.StreamFileUpload(strings.Join([]string{"v1", env.TemplateID, buildID.String(), "context.tar.gz"}, "/"), fileContent)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when uploading file to cloud storage: %s", err))

		err = fmt.Errorf("error when uploading file to cloud storage: %w", err)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	err = a.buildCache.Create(env.TemplateID, buildID, team.ID)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusConflict, fmt.Sprintf("there's already running build for %s", env.TemplateID))

		err = fmt.Errorf("build is already running build for %s", env.TemplateID)
		telemetry.ReportCriticalError(ctx, err)

		return nil
	}

	telemetry.ReportEvent(ctx, "started updating environment")

	if alias != "" {
		err = a.supabase.EnsureEnvAlias(ctx, alias, env.TemplateID)
		if err != nil {
			a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when inserting alias: %s", err))

			err = fmt.Errorf("error when inserting alias: %w", err)
			telemetry.ReportCriticalError(ctx, err)

			a.buildCache.Delete(env.TemplateID, buildID, team.ID)

			return nil
		} else {
			telemetry.ReportEvent(ctx, "inserted alias", attribute.String("env.alias", alias))
		}
	}

	startCmd := c.PostForm("startCmd")

	go func() {
		buildContext, childSpan := a.tracer.Start(
			trace.ContextWithSpanContext(context.Background(), span.SpanContext()),
			"background-build-env",
		)

		var status api.TemplateBuildStatus

		buildErr := a.buildEnv(
			buildContext,
			userID.String(),
			team.ID,
			env.TemplateID,
			buildID,
			dockerfile,
			startCmd,
			envKernelVersion,
			envFirecrackerVersion,
			properties,
			nomad.BuildConfig{
				VCpuCount:          cpuCount,
				MemoryMB:           ramMB,
				DiskSizeMB:         tier.DiskMB,
				KernelVersion:      schema.DefaultKernelVersion,
				FirecrackerVersion: schema.DefaultFirecrackerVersion,
			})

		if buildErr != nil {
			status = api.TemplateBuildStatusError

			errMsg := fmt.Errorf("error when building env: %w", buildErr)

			telemetry.ReportCriticalError(buildContext, errMsg)
		} else {
			status = api.TemplateBuildStatusReady

			telemetry.ReportEvent(buildContext, "created new environment", attribute.String("env.id", env.TemplateID))
		}

		if status == api.TemplateBuildStatusError && alias != "" {
			errMsg := a.supabase.DeleteNilEnvAlias(buildContext, alias)
			if errMsg != nil {
				err = fmt.Errorf("error when deleting alias: %w", errMsg)
				telemetry.ReportError(buildContext, err)
			} else {
				telemetry.ReportEvent(buildContext, "deleted alias", attribute.String("env.alias", alias))
			}
		} else if status == api.TemplateBuildStatusReady && alias != "" {
			errMsg := a.supabase.UpdateEnvAlias(buildContext, alias, env.TemplateID)
			if errMsg != nil {
				err = fmt.Errorf("error when updating alias: %w", errMsg)
				telemetry.ReportError(buildContext, err)
			} else {
				telemetry.ReportEvent(buildContext, "updated alias", attribute.String("env.alias", alias))
			}
		}

		cacheErr := a.buildCache.SetDone(env.TemplateID, buildID, status)
		if cacheErr != nil {
			errMsg := fmt.Errorf("error when setting build done in logs: %w", cacheErr)
			telemetry.ReportCriticalError(buildContext, errMsg)
		}

		childSpan.End()
	}()

	aliases := []string{}

	if alias != "" {
		aliases = append(aliases, alias)
	}

	a.logger.Infof("Built template %s with build id %s", env.TemplateID, buildID.String())

	return &api.Template{
		TemplateID: env.TemplateID,
		BuildID:    buildID.String(),
		Public:     false,
		Aliases:    &aliases,
	}
}

func (a *APIStore) buildEnv(
	ctx context.Context,
	userID string,
	teamID uuid.UUID,
	envID string,
	buildID uuid.UUID,
	dockerfile,
	startCmd,
	envKernelVersion,
	envFirecrackerVersion string,
	posthogProperties posthog.Properties,
	vmConfig nomad.BuildConfig,
) (err error) {
	childCtx, childSpan := a.tracer.Start(ctx, "build-env",
		trace.WithAttributes(
			attribute.String("env.id", envID),
			attribute.String("build.id", buildID.String()),
			attribute.String("env.team.id", teamID.String()),
		),
	)
	defer childSpan.End()

	startTime := time.Now()

	defer func() {
		a.posthog.CreateAnalyticsUserEvent(userID, teamID.String(), "built environment", posthogProperties.
			Set("environment", envID).
			Set("build_id", buildID).
			Set("duration", time.Since(startTime).String()).
			Set("success", err != nil),
		)
	}()

	diskSize, err := a.nomad.BuildEnvJob(
		a.tracer,
		childCtx,
		envID,
		envKernelVersion,
		envFirecrackerVersion,
		buildID.String(),
		startCmd,
		a.apiSecret,
		a.googleServiceAccountBase64,
		vmConfig,
	)
	if err != nil {
		err = fmt.Errorf("error when building env: %w", err)
		telemetry.ReportCriticalError(childCtx, err)

		return err
	}

	err = a.supabase.UpsertEnv(
		ctx,
		teamID,
		envID,
		buildID,
		dockerfile,
		vmConfig.VCpuCount,
		vmConfig.MemoryMB,
		vmConfig.DiskSizeMB,
		diskSize,
		schema.DefaultKernelVersion,
		schema.DefaultFirecrackerVersion,
	)
	if err != nil {
		err = fmt.Errorf("error when updating env: %w", err)
		telemetry.ReportCriticalError(childCtx, err)

		return err
	}

	return nil
}

func getCPUAndRAM(tierID, cpuStr, memoryMBStr string) (int64, int64, *api.APIError) {
	cpu := constants.DefaultTemplateCPU
	ramMB := constants.DefaultTemplateMemory

	// Check if team can customize the resources
	if (cpuStr != "" || memoryMBStr != "") && tierID == constants.BaseTierID {
		return 0, 0, &api.APIError{
			Err:       fmt.Errorf("team with tier %s can't customize resources", tierID),
			ClientMsg: "Team with this tier can't customize resources, don't specify cpu count or memory",
			Code:      http.StatusBadRequest,
		}
	}

	if cpuStr != "" {
		cpuInt, err := strconv.Atoi(cpuStr)
		if err != nil {
			return 0, 0, &api.APIError{
				Err:       fmt.Errorf("error when parsing customCPUs: %w", err),
				ClientMsg: "CPU count must be a number",
				Code:      http.StatusBadRequest,
			}
		}

		if cpuInt < constants.MinTemplateCPU || cpuInt > constants.MaxTemplateCPU {
			return 0, 0, &api.APIError{
				Err:       err,
				ClientMsg: fmt.Sprintf("CPU must be between %d and %d", constants.MinTemplateCPU, constants.MaxTemplateCPU),
				Code:      http.StatusBadRequest,
			}
		}

		cpu = cpuInt
	}

	if memoryMBStr != "" {
		memoryMBInt, err := strconv.Atoi(memoryMBStr)
		if err != nil {
			return 0, 0, &api.APIError{
				Err:       err,
				ClientMsg: "Memory must be a number",
				Code:      http.StatusBadRequest,
			}
		}

		if memoryMBInt < constants.MinTemplateMemory || memoryMBInt > constants.MaxTemplateMemory {
			return 0, 0, &api.APIError{
				Err:       err,
				ClientMsg: fmt.Sprintf("Memory must be between %d and %d", constants.MinTemplateMemory, constants.MaxTemplateMemory),
				Code:      http.StatusBadRequest,
			}
		}

		if memoryMBInt%2 != 0 {
			return 0, 0, &api.APIError{
				Err:       fmt.Errorf("customMemory must be divisible by 2"),
				ClientMsg: "Memory must be a divisible by 2",
				Code:      http.StatusBadRequest,
			}
		}

		ramMB = memoryMBInt
	}

	return int64(cpu), int64(ramMB), nil
}
