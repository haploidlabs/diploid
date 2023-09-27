package domain

import (
	"database/sql"
	"github.com/haploidlabs/diploid/internal/db"
)

// Environment represents a project environment.
type Environment struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ToDB converts an Environment to a db.ProjectEnvironment.
func (e *Environment) ToDB() *db.Environment {
	return &db.Environment{
		ID:   e.ID,
		Name: e.Name,
		Description: sql.NullString{
			String: e.Description,
		},
	}
}

// EnvironmentFromDB converts a db.ProjectEnvironment to an Environment.
func EnvironmentFromDB(e *db.Environment) *Environment {
	return &Environment{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description.String,
	}
}

type CreateEnvironmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateEnvironmentRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
