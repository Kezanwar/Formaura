package routes

import (
	"formaura/cmd/api/handlers"
	"formaura/pkg/output"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	output.MakeRoute(r, "/", handlers.GetUsers).Methods("GET", "OPTIONS")
}
