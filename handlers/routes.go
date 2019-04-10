package handlers

import (
	"net/http"

	"metahub/handlers/registry"
)

// Register registers handlers/routers
func Register() {
	http.Handle("/v2/", registry.NewRouter())
}
