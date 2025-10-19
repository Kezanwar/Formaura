package main

import (
	"context"
	"formaura/cmd/api/handlers"
	"formaura/cmd/api/routes"
	user_memory_cache "formaura/pkg/cache/user_memory"
	"formaura/pkg/email"
	"formaura/pkg/middleware"
	form_repo "formaura/pkg/repositories/form"
	user_repo "formaura/pkg/repositories/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewAPI(ctx context.Context, pool *pgxpool.Pool, client *http.Client) (*http.Server, error) {

	TWO_HOURS := 2 * time.Hour

	emailClient, err := email.NewClient()

	if err != nil {
		log.Fatalf("Email client failed to init: %v", err)
	}

	//memory cache
	userCache := user_memory_cache.New(TWO_HOURS)

	//repositories
	userRepo := user_repo.NewUserRepo(pool)
	formRepo := form_repo.NewFormRepo(pool)

	//handlers
	authHandlers := handlers.NewAuthHandler(userRepo, userCache, emailClient)
	formHandlers := handlers.NewFormHandler(formRepo, userCache, emailClient)

	authFresh := middleware.AuthAlwaysFreshMiddleware(userRepo, userCache)
	authCached := middleware.AuthCachedMiddleware(userRepo, userCache)

	//router
	r := mux.NewRouter()
	r.Use(middleware.Cors)
	api := r.PathPrefix("/api").Subrouter()

	//apply routes
	routes.Register(
		//main router
		api,
		//handlers
		authHandlers,
		formHandlers,
		//middleware
		authFresh,
		authCached,
	)

	return &http.Server{
		Addr:    PORT,
		Handler: r,
	}, nil
}
