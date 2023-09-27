package domain

import (
	"database/sql"
	"github.com/haploidlabs/diploid/internal/db"
	"time"
)

// Project represents a Diploid project.
type Project struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int64     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

// ToDB converts a Project to a db.Project.
func (p *Project) ToDB() *db.Project {
	return &db.Project{
		ID:   p.ID,
		Name: p.Name,
		Description: sql.NullString{
			String: p.Description,
			Valid:  true,
		},
		CreatedBy: p.CreatedBy,
		CreatedAt: sql.NullTime{
			Time:  p.CreatedAt,
			Valid: true,
		},
	}
}

// ProjectFromDB converts a db.Project to a Project.
func ProjectFromDB(p *db.Project) *Project {
	return &Project{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description.String,
		CreatedBy:   p.CreatedBy,
		CreatedAt:   p.CreatedAt.Time,
	}
}

// CreateProjectRequest represents a request to create a project.
type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
