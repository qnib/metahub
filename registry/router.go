package registry

import (
	"fmt"
	"metahub/machinetypes"
	"net/http"

	"metahub/environment"

	"github.com/docker/distribution/reference"
	"github.com/gorilla/mux"
)

type registry struct {
	env environment.Environment
}

// NewRouter returns a router for the registry API endpoints
func NewRouter(env environment.Environment, pathPrefix string) http.Handler {
	reg := registry{
		env: env,
	}
	router := mux.NewRouter()
	router.Use(machinetypes.Middleware())
	router.HandleFunc(pathPrefix+"/{image}/manifests/{reference}", reg.manifestsHandler).Methods("GET")
	router.HandleFunc(pathPrefix+"/{repo}/{image}/manifests/{reference}", reg.manifestsHandler).Methods("GET")
	router.HandleFunc(pathPrefix+"/{image}/blobs/{reference}", reg.blobsHandler).Methods("GET")
	router.HandleFunc(pathPrefix+"/{repo}/{image}/blobs/{reference}", reg.blobsHandler).Methods("GET")
	router.HandleFunc(pathPrefix+"/", reg.baseHandler).Methods("GET")
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
