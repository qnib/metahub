package registry

import (
	"fmt"
	"net/http"

	"metahub/environment"

	basicAuth "github.com/99designs/basicauth-go"
	manifestListSchema "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema "github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/reference"
	"github.com/gorilla/mux"
)

type registry struct {
	env environment.Environment
}

var authMiddleware func(http.Handler) http.Handler

func init() {
	_ = manifestListSchema.SchemaVersion
	_ = manifestSchema.SchemaVersion

	authMiddleware = basicAuth.New("MetaHub", map[string][]string{
		"test1":  {"test1"},
		"test2":  {"test2"},
		"test3":  {"test3"},
		"test12": {"test12"},
	})
}

// NewRouter returns a router for the registry API endpoints
func NewRouter(env environment.Environment) http.Handler {
	reg := registry{
		env: env,
	}
	router := mux.NewRouter()
	router.Use(authMiddleware)
	router.HandleFunc("/v2/{image}/manifests/{reference}", reg.manifestsHandler).Methods("GET")
	router.HandleFunc("/v2/{repo}/{image}/manifests/{reference}", reg.manifestsHandler).Methods("GET")
	router.HandleFunc("/v2/{image}/blobs/{reference}", reg.blobsHandler).Methods("GET")
	router.HandleFunc("/v2/{repo}/{image}/blobs/{reference}", reg.blobsHandler).Methods("GET")
	router.HandleFunc("/v2/", reg.baseHandler).Methods("GET")
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
