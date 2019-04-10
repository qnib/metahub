package registry

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
)

func blobsHandler(w http.ResponseWriter, r *http.Request) {

	_ = r.Context()
	vars := mux.Vars(r)

	// get image repository
	repositoryName, err := getRepository(r)
	if err != nil {
		log.Printf("error parsing image repository: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = repositoryName

	// get blob digest
	digest, err := digest.Parse(vars["reference"])
	if err != nil {
		log.Printf("error parsing blob reference: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = digest

}
