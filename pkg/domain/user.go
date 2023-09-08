package domain

import (
	"database/sql"
	"github.com/haploidlabs/diploid/internal/db"
	"time"
)

// User represents a Diploid user.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// ToDB converts a User to a db.User.
func (u *User) ToDB() *db.User {
	return &db.User{
		ID:    int64(u.ID),
		Name:  u.Name,
		Email: u.Email,
		CreatedAt: sql.NullTime{
			Time: u.CreatedAt,
		},
	}
}

// UserFromDB converts a db.User to a User.
func UserFromDB(u *db.User) *User {
	return &User{
		ID:        int(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
	}
}
