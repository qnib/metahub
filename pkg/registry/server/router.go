package server

import (
	"fmt"
	"metahub/pkg/daemon"
	"metahub/pkg/machinetypes"
	"net/http"

	"github.com/docker/distribution/reference"
	"github.com/gorilla/mux"
)

// NewRouter returns a router for the registry API endpoints
func NewRouter(service daemon.Service, pathPrefix string) http.Handler {
	router := mux.NewRouter()
	router.Use(machinetypes.AuthMiddleware(service))
	router.Handle(pathPrefix+"/{image}/manifests/{reference}", getRegistryHandler(service)).Methods("GET")
	router.Handle(pathPrefix+"/{repo}/{image}/manifests/{reference}", getRegistryHandler(service)).Methods("GET")
	router.Handle(pathPrefix+"/{image}/blobs/{reference}", getBlobsHandler(service)).Methods("GET")
	router.Handle(pathPrefix+"/{repo}/{image}/blobs/{reference}", getBlobsHandler(service)).Methods("GET")
	router.Handle(pathPrefix+"/", getBaseHandler(service)).Methods("GET")
	return router
}

func getRepository(r *http.Request) (reference.Named, error) {
	vars := mux.Vars(r)
	image := vars["image"]
	repo := vars["repo"]
	if repo == "" {
		repo = "library"
	}
	name := repo + "/" + image
	n, err := reference.WithName(name)
	if err != nil {
		return nil, fmt.Errorf("error parsing repository name: %v", err)
	}
	return n, nil
}
