package api

import (
	"encoding/json"
	"github.com/haploidlabs/diploid/pkg/domain"
	"log/slog"
	"net/http"
)

func UserID(r *http.Request) int64 {
	return r.Context().Value("userId").(int64)
}

func (api *API) User(w http.ResponseWriter, r *http.Request) *domain.User {
	u, err := api.DB.GetUserByID(r.Context(), r.Context().Value("userId").(int64))
	if err != nil {
		WriteError(w, http.StatusNotFound, domain.ErrUserNotFound)
		return nil
	}
	return domain.UserFromDB(&u)
}

func WriteStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		slog.Error("failed to write response", err)
		return
	}
}

func WriteError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		WriteJSON(w, code, struct{}{})
		return
	}
	WriteJSON(w, code, map[string]string{
		"error": err.Error(),
	})
}

func DecodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
	}
	return err
}

func Int64FromString(w http.ResponseWriter, s string) int64 {
	var i int64
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		WriteError(w, http.StatusBadRequest, domain.ErrBadRequest)
		return 0
	}
	return i
}
