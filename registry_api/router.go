package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/v2/{image}/manifests/{reference}", manifestHandler).Methods("GET")
	r.HandleFunc("/v2/{repo}/{image}/manifests/{reference}", manifestHandler).Methods("GET")
	r.HandleFunc("/v2/", baseHandler).Methods("GET")
	http.Handle("/v2/", r)
}
