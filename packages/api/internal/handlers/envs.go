package handlers

import (
	"fmt"
	"github.com/e2b-dev/api/packages/api/internal/constants"
	"github.com/e2b-dev/api/packages/api/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (a *APIStore) PostEnvs(
	c *gin.Context,
) {
	ctx := c.Request.Context()

	fileContent, fileHandler, err := c.Request.FormFile("buildContext")
	if err != nil {
		formErr := fmt.Errorf("error when parsing form data: %w", err)
		ReportCriticalError(ctx, formErr)
		return
	}
	defer fileContent.Close()

	// Check if file is a tar.gz file
	if !strings.HasSuffix(fileHandler.Filename, ".tar.gz") {
		a.sendAPIStoreError(c, http.StatusBadRequest, "Build context must be a tar.gz file")

		return
	}

	// Upload file to cloud storage
	err = a.uploadFile(fileHandler.Filename, fileContent)
	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when uploading file: %s", err))

		return
	}

	// Not implemented yet
	envID := c.PostForm("envID")
	if envID != "" {
		a.sendAPIStoreError(c, http.StatusNotImplemented, "Updating envs is not implemented yet")

		return
	}

	// Prepare info for new env
	envID = utils.GenerateID()
	userID := c.Value(constants.UserIDContextKey).(string)
	team, err := a.supabase.GetDefaultTeamFromUserID(userID)

	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when getting default team: %s", err))

		return
	}

	// Save env to database
	newEnv, err := a.supabase.CreateEnv(envID, team.ID, c.PostForm("dockerfile"))
	if err != nil {
		a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Error when creating env: %s", err))

		return
	}

	c.JSON(http.StatusOK, newEnv)
}
