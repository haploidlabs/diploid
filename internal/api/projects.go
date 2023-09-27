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

func (api *API) HandleCreateProject(w http.ResponseWriter, r *http.Request) {
	uid := UserID(r)

	// Decode request body
	var dto domain.CreateProjectRequest
	err := DecodeBody(w, r, &dto)
	if err != nil {
		return
	}

	// Validate request body
	val := valgo.Is(valgo.String(dto.Name, "name").Not().Blank().OfLengthBetween(3, 30)).Is(valgo.String(dto.Description, "description").MaxLength(255))
	if !val.Valid() {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}

	// Create project
	p, err := api.DB.CreateProject(r.Context(), db.CreateProjectParams{
		Name: dto.Name,
		Description: sql.NullString{
			String: dto.Description,
			Valid:  true,
		},
		CreatedBy: uid,
	})

	if err != nil {
		slog.Error("failed to create project", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	// Write response
	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"project": domain.ProjectFromDB(&p),
	})
}

func (api *API) HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	uid := UserID(r)

	// Get projects
	p, err := api.DB.GetProjectsByUser(r.Context(), uid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		slog.Error("failed to get projects", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	// Convert to domain.Project
	projects := make([]*domain.Project, len(p))
	for i, project := range p {
		projects[i] = domain.ProjectFromDB(&project)
	}

	// Write response
	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"projects": projects,
	})
}

func (api *API) HandleUpdateProject(w http.ResponseWriter, r *http.Request) {
	uid := UserID(r)

	rawProjectId := chi.URLParam(r, "projectID")
	if rawProjectId == "" {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	projectID := Int64FromString(w, rawProjectId)
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

	// Decode request body
	var dto domain.UpdateProjectRequest
	err = DecodeBody(w, r, &dto)
	if err != nil {
		return
	}

	// Validate request body
	if dto.Name != "" {
		val := valgo.Is(valgo.String(dto.Name, "name").Not().Blank().OfLengthBetween(3, 30))
		if !val.Valid() {
			WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
			return
		}

		if dto.Name == p.Name {
			// Check if project with name exists
			_, err = api.DB.GetProjectByName(r.Context(), db.GetProjectByNameParams{
				Name:      dto.Name,
				CreatedBy: uid,
			})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				slog.Error("failed to get project by name", err)
				WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
				return
			}
			if errors.Is(err, sql.ErrNoRows) {
				WriteError(w, http.StatusConflict, domain.ErrProjectWithNameExists)
				return
			}
		}
		p.Name = dto.Name
	}

	if dto.Description != "" {
		val := valgo.Is(valgo.String(dto.Description, "description").MaxLength(255))
		if !val.Valid() {
			WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
			return
		}
		p.Description = sql.NullString{
			String: dto.Description,
			Valid:  true,
		}
	}

	// Update project
	_, err = api.DB.UpdateProject(r.Context(), db.UpdateProjectParams{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	})
	if err != nil {
		slog.Error("failed to update project", err)
		WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
		return
	}

	// Write response
	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"project": domain.ProjectFromDB(&p),
	})
}
