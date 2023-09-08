package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/haploidlabs/diploid/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"time"
)

func (api *API) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto domain.LoginRequest
		err := DecodeBody(w, r, &dto)
		if err != nil {
			return
		}

		// Get user by email
		ctx, ccl := context.WithTimeout(r.Context(), 5*time.Second)
		defer ccl()
		u, err := api.DB.GetUserByEmail(ctx, dto.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				WriteError(w, http.StatusNotFound, domain.ErrUserNotFound)
				return
			}
			slog.Error("failed to get user by email", err)
			WriteError(w, http.StatusInternalServerError, domain.ErrInternal)
			return
		}

		// Compare password
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(dto.Password))
		if err != nil {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}

		// Generate JWT
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": u.ID,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
		}).SignedString([]byte(api.JWTSecret))

		// Write response
		WriteJSON(w, http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
