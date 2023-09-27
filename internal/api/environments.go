package api

import (
	"database/sql"
	"errors"
	"github.com/cohesivestack/valgo"
	"github.com/go-chi/chi/v5"
	"github.com/haploidlabs/diploid/internal/db"
	"github.com/haploidlabs/diploid/pkg/domain"
	"log/slog"
	"net/http"
)

func (api *API) HandleCreateEnvironment(w http.ResponseWriter, r *http.Request) {
	uid := UserID(r)

	// check project param
	rawProjectID := chi.URLParam(r, "projectID")
	if rawProjectID == "" {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	projectID := Int64FromString(w, rawProjectID)
	if projectID == 0 {
		return
	}

	// check if project exists
	p, err := api.DB.GetProjectByID(r.Context(), projectID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		WriteError(w, http.StatusNotFound, domain.ErrNotFound)
		return
	}
	if err != nil {
		slog.Error("failed to get project by id", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	// check if user is member of project
	if p.CreatedBy != uid {
		WriteError(w, http.StatusForbidden, domain.ErrUnauthorized)
		return
	}

	var dto domain.CreateEnvironmentRequest
	err = DecodeBody(w, r, &dto)
	if err != nil {
		return
	}

	// validate dto
	val := valgo.Is(valgo.String(dto.Name, "name").Not().Blank().OfLengthBetween(3, 50)).Is(valgo.String(dto.Description, "description").MaxLength(255))
	if !val.Valid() {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}

	// Check if environment with name exists
	_, err = api.DB.GetEnvironmentByName(r.Context(), db.GetEnvironmentByNameParams{
		Name:      dto.Name,
		ProjectID: projectID,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		slog.Error("failed to get environment by name", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		WriteError(w, http.StatusConflict, domain.ErrEnvironmentWithNameExists)
		return
	}

	// Create environment
	e, err := api.DB.CreateEnvironment(r.Context(), db.CreateEnvironmentParams{
		ProjectID: projectID,
		Name:      dto.Name,
		Description: sql.NullString{
			String: dto.Description,
		},
	})

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"environment": domain.EnvironmentFromDB(&e),
	})
}

func (api *API) HandleGetEnvironments(w http.ResponseWriter, r *http.Request) {
	uid := UserID(r)

	// check project param
	rawProjectID := chi.URLParam(r, "projectID")
	if rawProjectID == "" {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	projectID := Int64FromString(w, rawProjectID)
	if projectID == 0 {
		return
	}

	// check if project exists
	p, err := api.DB.GetProjectByID(r.Context(), projectID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		WriteError(w, http.StatusNotFound, domain.ErrNotFound)
		return
	}
	if err != nil {
		slog.Error("failed to get project by id", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	// check if user is member of project
	if p.CreatedBy != uid {
		WriteError(w, http.StatusForbidden, domain.ErrUnauthorized)
		return
	}

	// Get environments
	environments, err := api.DB.GetEnvironmentsForProject(r.Context(), projectID)
	if err != nil {
		slog.Error("failed to get environments for project", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	// convert to domain.Environment
	envs := make([]*domain.Environment, len(environments))
	for i, e := range environments {
		envs[i] = domain.EnvironmentFromDB(&e)
	}
	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"environments": envs,
	})
}

func (api *API) HandleUpdateEnvironment(w http.ResponseWriter, r *http.Request) {
	uid := UserID(r)

	// check project param
	rawProjectID := chi.URLParam(r, "projectID")
	if rawProjectID == "" {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	projectID := Int64FromString(w, rawProjectID)
	if projectID == 0 {
		return
	}

	// check if project exists
	p, err := api.DB.GetProjectByID(r.Context(), projectID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		WriteError(w, http.StatusNotFound, domain.ErrNotFound)
		return
	}
	// check if user is member of project
	if p.CreatedBy != uid {
		WriteError(w, http.StatusForbidden, domain.ErrUnauthorized)
		return
	}

	// check environment param
	rawEnvironmentID := chi.URLParam(r, "environmentID")
	if rawEnvironmentID == "" {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	environmentID := Int64FromString(w, rawEnvironmentID)
	if environmentID == 0 {
		return
	}

	// check if environment exists
	e, err := api.DB.GetEnvironmentById(r.Context(), environmentID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		WriteError(w, http.StatusNotFound, domain.ErrNotFound)
		return
	}
	if err != nil {
		slog.Error("failed to get environment by id", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	var dto domain.UpdateEnvironmentRequest
	err = DecodeBody(w, r, &dto)
	if err != nil {
		return
	}

	if dto.Name != "" {
		val := valgo.Is(valgo.String(dto.Name, "name").Not().Blank().OfLengthBetween(3, 50))
		if !val.Valid() {
			WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
			return
		}

		if dto.Name != e.Name {
			// Check if environment with name exists
			_, err = api.DB.GetEnvironmentByName(r.Context(), db.GetEnvironmentByNameParams{
				Name:      dto.Name,
				ProjectID: projectID,
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				slog.Error("failed to get environment by name", err)
				WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
				return
			}
			if !errors.Is(err, sql.ErrNoRows) {
				WriteError(w, http.StatusConflict, domain.ErrEnvironmentWithNameExists)
				return
			}
		}
		e.Name = dto.Name
	}

	if dto.Description != "" {
		val := valgo.Is(valgo.String(dto.Description, "description").MaxLength(255))
		if !val.Valid() {
			WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
			return
		}
		e.Description = sql.NullString{
			String: dto.Description,
		}
	}

	// Update environment
	_, err = api.DB.UpdateEnvironment(r.Context(), db.UpdateEnvironmentParams{
		ID:          environmentID,
		Name:        e.Name,
		Description: e.Description,
	})

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"environment": domain.EnvironmentFromDB(&e),
	})
}
