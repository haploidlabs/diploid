package main

import (
	"context"
	"database/sql"
	"github.com/haploidlabs/diploid/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/haploidlabs/diploid/internal/api"
	"github.com/haploidlabs/diploid/internal/config"
	"github.com/haploidlabs/diploid/internal/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.LoadConfig()

	// Create Database Client
	dbClient, err := sql.Open("sqlite3", cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := dbClient.Close()
		if err != nil {
			slog.Error("failed to close database connection", err)
		}
	}()

	// Create Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(cfg.CorsAllowOrigins, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Docker Client
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("failed to create docker client: %v", err)
	}

	sqlcDB := db.New(dbClient)

	// Seed database
	seed(sqlcDB)

	// Serve Vue app static files
	r.Get("/_nuxt/*", serveStatic("./web/dist/_nuxt"))

	// Catch-all route for client-side routing
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/dist/index.html")
	})

	// Create and start API server
	s := api.New(api.Options{
		Router:         r,
		AllowedOrigins: strings.Split(cfg.CorsAllowOrigins, ","),
		DB:             sqlcDB,
		JWTSecret:      cfg.JWTSecret,
		DockerClient:   dockerCli,
	})
	s.Start(cfg.BindAddress)
}

func serveStatic(basePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/_nuxt")
		fullPath := filepath.Join(basePath, path)
		ext := filepath.Ext(fullPath)

		switch ext {
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".svg":
			w.Header().Set("Content-Type", "image/svg+xml")
			// Add more cases for other file types if needed
		}

		http.ServeFile(w, r, fullPath)
	}
}

func seed(sqlcDB *db.Queries) {
	seedUser(sqlcDB)
}

func seedUser(sqlcDB *db.Queries) {
	ctx, ccl := context.WithTimeout(context.Background(), 10*time.Second)
	defer ccl()

	// Check if users exist
	amount, err := sqlcDB.CountUsers(ctx)
	if err != nil {
		log.Fatalf("failed to request users: %v", err)
	}
	if amount > 0 {
		return
	}

	// Create seed admin user
	pw, err := bcrypt.GenerateFromPassword([]byte("admin1234"), bcrypt.DefaultCost)
	_, err = sqlcDB.CreateUser(ctx, db.CreateUserParams{
		Name:     "Admin",
		Email:    "admin@diploid.dev",
		Password: string(pw),
		Role:     domain.UserRoleAdmin,
	})

	if err != nil {
		log.Fatalf("failed to create seed user: %v", err)
	}
}
