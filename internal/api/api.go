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

	v1.Post("/auth/login", api.HandleLogin)

	// Authenticated routes
	v1Auth := v1.Group(nil)
	v1Auth.Use(api.Auth)
	v1Auth.Get("/projects", api.HandleGetProjects)
	v1Auth.Post("/projects", api.HandleCreateProject)
	v1Auth.Put("/projects/{projectID}", api.HandleUpdateProject)
	v1Auth.Get("/projects/{projectID}/environments", api.HandleGetEnvironments)
	v1Auth.Post("/projects/{projectID}/environments", api.HandleCreateEnvironment)
	v1Auth.Put("/projects/{projectID}/environments/{environmentID}", api.HandleUpdateEnvironment)

	api.r.Mount("/v1", v1)
}

func (api *API) Start(addr string) {
	log.Printf("server running on %+v", addr)
	err := http.ListenAndServe(addr, api.r)
	if err != nil {
		log.Fatalf("failed to serve application: %+v", err)
	}
}
