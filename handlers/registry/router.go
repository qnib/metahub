package registry

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns a router for the registry API endpoints
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/v2/{image}/manifests/{reference}", manifestsHandler).Methods("GET")
	r.HandleFunc("/v2/{repo}/{image}/manifests/{reference}", manifestsHandler).Methods("GET")
	r.HandleFunc("/v2/{image}/blobs/{reference}", blobsHandler).Methods("GET")
	r.HandleFunc("/v2/{repo}/{image}/blobs/{reference}", blobsHandler).Methods("GET")
	r.HandleFunc("/v2/", baseHandler).Methods("GET")
	return r
}
