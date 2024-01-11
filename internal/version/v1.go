package v1

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/portfolio-builder/internal/routes/v1"
)

func Routes(router *mux.Router) {
	// handle versioning
	r := router.PathPrefix("/").Subrouter()

	// Register routes
	routes.Routes(r)
}
