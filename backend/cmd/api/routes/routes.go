package routes

import (
	"formaura/cmd/api/handlers"
	"formaura/pkg/middleware"
	"formaura/pkg/output"

	"github.com/gorilla/mux"
)

func Register(
	//router
	r *mux.Router,
	//handlers
	authHandler *handlers.AuthHandler,

	//middlewares
	authFresh middleware.Middleware,
	authCached middleware.Middleware) {

	output.MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandler, authCached)
	})

}
