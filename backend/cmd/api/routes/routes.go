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
	authHandlers *handlers.AuthHandler,
	formHandlers *handlers.FormHandler,

	//middlewares
	authFresh middleware.Middleware,
	authCached middleware.Middleware) {

	output.MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandlers, authCached)
	})
	output.MakeSubRouter(r, "/form", func(sr *mux.Router) {
		FormRoutes(sr, formHandlers, authCached)
	})

}
