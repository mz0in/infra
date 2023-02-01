package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/devbookhq/devbook-api/packages/api/internal/api"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (a *APIStore) GetSessions(
	c *gin.Context,
	params api.GetSessionsParams,
) {
	ctx := c.Request.Context()

	_, admin, keyErr := a.validateAPIKey(&params.ApiKey)
	if keyErr != nil {
		errMsg := fmt.Errorf("error with API key: %+v", keyErr)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, http.StatusUnauthorized, "Error with API token")
		return
	}
	ReportEvent(ctx, "validated API key")

	if !admin {
		ReportError(ctx, fmt.Errorf("Unauthorized"))
		a.sendAPIStoreError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	sessions, err := a.nomad.GetSessions()
	if err != nil {
		errMsg := fmt.Errorf("error when listing sessions: %v", err)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, err.Code, err.ClientMsg)
		return
	}
	ReportEvent(ctx, "listed sessions")

	c.JSON(http.StatusOK, sessions)
}

func (a *APIStore) PostSessions(
	c *gin.Context,
	params api.PostSessionsParams,
) {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)

	var newSession api.PostSessionsJSONRequestBody
	if err := c.Bind(&newSession); err != nil {
		errMsg := fmt.Errorf("error when parsing request: %s", err)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, http.StatusBadRequest, fmt.Sprintf("Error when parsing request: %s", err))
		return
	}
	ReportEvent(ctx, "parsed request")

	// The default option in the openapi does not automatically populate JSON field with the default value
	if newSession.EditEnabled == nil {
		editEnabled := false
		newSession.EditEnabled = &editEnabled
	}

	span.SetAttributes(attribute.Bool("session.edit_enabled", *newSession.EditEnabled))

	if *newSession.EditEnabled {
		userID, admin, keyErr := a.validateAPIKey(params.ApiKey)
		if keyErr != nil {
			errMsg := fmt.Errorf("error with API key: %+v", keyErr)
			ReportCriticalError(ctx, errMsg)
			a.sendAPIStoreError(c, http.StatusUnauthorized, "Error with API token")
			return
		}
		ReportEvent(ctx, "validated API key")

		if !admin {
			owner, err := a.isOwner(newSession.CodeSnippetID, userID)
			if err != nil {
				ReportError(ctx, fmt.Errorf("error getting user data from Supabase: %+v", err))
				a.sendAPIStoreError(c, http.StatusInternalServerError, fmt.Sprintf("Cannot retrieve data: %s", err))
				return
			}
			if !owner {
				a.sendAPIStoreError(c, http.StatusUnauthorized, "Unauthorized")
				return
			}
		}

		existingSession, err := a.sessionsCache.FindEditSession(newSession.CodeSnippetID)
		if err != nil {
			ReportEvent(ctx, "no existing edit session found")
		} else {
			ReportEvent(ctx, "found existing edit session")
			c.JSON(http.StatusCreated, &existingSession)
			return
		}
	}

	session, err := a.nomad.CreateSession(a.tracer, ctx, &newSession)
	if err != nil {
		errMsg := fmt.Errorf("error when creating: %v", err)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, err.Code, err.ClientMsg)
		return
	}
	ReportEvent(ctx, "created session")

	if *newSession.EditEnabled {
		// We check for the edit session again because we didn't want to lock for the whole duration of this function.
		// If we find an existing session now we just discard the one we created and everything will work.
		existingSession, err := a.sessionsCache.FindEditSession(newSession.CodeSnippetID)
		if err == nil {
			fmt.Printf("Found another edit session after we created a new editing session. Returning the other session.")

			delErr := a.nomad.DeleteSession(session.SessionID, true)
			if delErr != nil {
				errMsg := fmt.Errorf("redundant session couldn't be deleted: %v", delErr)
				ReportError(ctx, errMsg)
			} else {
				ReportEvent(ctx, "deleted redundant session")
			}

			SetAttributes(ctx,
				attribute.String("session_id", existingSession.SessionID),
			)

			c.JSON(http.StatusCreated, &existingSession)
			return
		}
	}

	go func(ctx context.Context) {
		if err := a.sessionsCache.Add(session); err != nil {
			errMsg := fmt.Errorf("error when adding session to cache: %v", err)
			ReportError(ctx, errMsg)

			delErr := a.nomad.DeleteSession(session.SessionID, true)
			if delErr != nil {
				errMsg := fmt.Errorf("couldn't delete session that couldn't be added to cache: %v", delErr)
				ReportError(ctx, errMsg)
			} else {
				ReportEvent(ctx, "deleted session that couldn't be added to cache")
			}
			return
		}
	}(ctx)

	SetAttributes(ctx,
		attribute.String("session_id", session.SessionID),
	)

	c.JSON(http.StatusCreated, &session)
}

func (a *APIStore) DeleteSessionsSessionID(
	c *gin.Context,
	sessionID string,
	params api.DeleteSessionsSessionIDParams,
) {
	ctx := c.Request.Context()

	SetAttributes(ctx,
		attribute.String("session_id", sessionID),
	)

	_, admin, keyErr := a.validateAPIKey(&params.ApiKey)
	if keyErr != nil {
		errMsg := fmt.Errorf("error with API key: %+v", keyErr)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, http.StatusUnauthorized, "Error with API token")
		return
	}

	if !admin {
		ReportError(ctx, fmt.Errorf("Unauthorized"))
		a.sendAPIStoreError(c, http.StatusUnauthorized, "Users cannot terminate sessions")
		return
	}

	// TODO: Delete session from cache

	err := a.nomad.DeleteSession(sessionID, true)
	if err != nil {
		errMsg := fmt.Errorf("error when deleting session: %v", err)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, err.Code, err.ClientMsg)
		return
	}
	ReportEvent(ctx, "deleted session")

	c.Status(http.StatusNoContent)
}

func (a *APIStore) PostSessionsSessionIDRefresh(
	c *gin.Context,
	sessionID string,
	params api.PostSessionsSessionIDRefreshParams,
) {
	ctx := c.Request.Context()

	SetAttributes(ctx,
		attribute.String("session_id", sessionID),
	)

	// TODO: Require auth for refreshing edit sessions

	err := a.sessionsCache.Refresh(sessionID)
	if err != nil {
		errMsg := fmt.Errorf("error when refreshing session: %v", err)
		ReportCriticalError(ctx, errMsg)
		a.sendAPIStoreError(c, http.StatusNotFound, fmt.Sprintf("Error refreshing session - session '%s' was not found", sessionID))
		return
	}
	ReportEvent(ctx, "refreshed session")

	c.Status(http.StatusNoContent)
}
