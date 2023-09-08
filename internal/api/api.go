package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/haploidlabs/diploid/internal/db"
)

type Options struct {
	Router         *chi.Mux
	AllowedOrigins []string
	DB             *db.Queries
	JWTSecret      string
}

type API struct {
	r         *chi.Mux
	DB        *db.Queries
	JWTSecret string
}

func New(opts Options) *API {
	s := &API{
		r:         opts.Router,
		DB:        opts.DB,
		JWTSecret: opts.JWTSecret,
	}
	s.registerRoutes()
	return s
}

func (api *API) registerRoutes() {
	// frontend (web)
	api.r.Handle("/", http.FileServer(http.Dir("./web/build")))

	// API Routes
	v1 := chi.NewRouter()

	// users
	v1Auth := v1.Group(nil)
	v1Auth.Use(api.Auth)
	v1.Post("/auth/login", api.HandleLogin())
	// v1.Post("/users", api.HandleCreateUser)
	// v1.Get("/users/@me", api.HandleGetUser)

	api.r.Mount("/v1", v1)
}

func (api *API) Start(addr string) {
	log.Printf("server running on %+v", addr)
	err := http.ListenAndServe(addr, api.r)
	if err != nil {
		log.Fatalf("failed to serve application: %+v", err)
	}
}
