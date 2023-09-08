package api

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/haploidlabs/diploid/pkg/domain"
	"net/http"
	"strings"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			WriteError(w, http.StatusBadRequest, domain.ErrInvalidAuthorizationHeader)
			return
		}
		rawToken := parts[1]

		t, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, domain.ErrInvalidToken
			}
			return []byte(api.JWTSecret), nil
		})
		if err != nil {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", t.Claims.(jwt.MapClaims)["sub"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
