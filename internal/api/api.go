package api

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
	"github.com/haploidlabs/diploid/internal/db"
	"github.com/haploidlabs/diploid/pkg/domain"
	"log"
	"net/http"
)

type Options struct {
	Router         *chi.Mux
	AllowedOrigins []string
	DB             *db.Queries
	JWTSecret      string
	DockerClient   *client.Client
}

type API struct {
	r            *chi.Mux
	DB           *db.Queries
	JWTSecret    string
	DockerClient *client.Client
}

func New(opts Options) *API {
	s := &API{
		r:            opts.Router,
		DB:           opts.DB,
		JWTSecret:    opts.JWTSecret,
		DockerClient: opts.DockerClient,
	}
	s.registerRoutes()
	return s
}

func (api *API) registerRoutes() {
	// API Routes
	v1 := chi.NewRouter()

	v1.Post("/auth/login", api.HandleLogin)

	// Authenticated routes
	v1Auth := v1.Group(nil)
	v1Auth.Use(api.Auth)
	v1Auth.Get("/projects", api.HandleGetProjects)
	v1Auth.Post("/projects", api.HandleCreateProject)
	v1Auth.Put("/projects/{projectID}", api.HandleUpdateProject)
	v1Auth.Delete("/projects/{projectID}", api.HandleDeleteProject)
	v1Auth.Get("/projects/{projectID}/environments", api.HandleGetEnvironments)
	v1Auth.Post("/projects/{projectID}/environments", api.HandleCreateEnvironment)
	v1Auth.Put("/projects/{projectID}/environments/{environmentID}", api.HandleUpdateEnvironment)

	// TEst
	v1Auth.Get("/containers", func(w http.ResponseWriter, r *http.Request) {
		cs, err := api.DockerClient.ContainerList(r.Context(), types.ContainerListOptions{
			All: true,
		})
		if err != nil {
			panic(err)
		}
		containers := make([]domain.Container, len(cs))
		for index, container := range cs {
			// Status
			var status domain.ContainerStatus
			if container.State == "running" {
				status = domain.ContainerStatusRunning
			} else {
				status = domain.ContainerStatusStopped
			}

			// Ports
			ports := make([]domain.ContainerPort, 0)
			for _, port := range container.Ports {
				if port.PublicPort == 0 || port.PrivatePort == 0 || port.IP != "0.0.0.0" {
					continue
				}
				ports = append(ports, domain.ContainerPort{
					PrivatePort: int64(port.PrivatePort),
					PublicPort:  int64(port.PublicPort),
					Type:        port.Type,
				})
			}

			// Volumes
			volumes := make([]domain.ContainerVolume, len(container.Mounts))
			for index, volume := range container.Mounts {
				volumes[index] = domain.ContainerVolume{
					Host:      volume.Source,
					Container: volume.Destination,
				}
			}

			containers[index] = domain.Container{
				ID:      container.ID,
				Name:    container.Names[0],
				Image:   container.Image,
				ImageID: container.ImageID,
				Status:  status,
				Ports:   ports,
				Volumes: volumes,
			}
		}
		WriteJSON(w, http.StatusOK, map[string]interface{}{
			"containers": containers,
		})
		WriteStatus(w, http.StatusOK)
	})

	api.r.Mount("/api/v1", v1)
}

func (api *API) Start(addr string) {
	log.Printf("server running on %+v", addr)
	err := http.ListenAndServe(addr, api.r)
	if err != nil {
		log.Fatalf("failed to serve application: %+v", err)
	}
}
