package registry

import (
	"fmt"
	"net/http"

	ref "github.com/docker/distribution/reference"
	"github.com/gorilla/mux"
)

func getRepository(r *http.Request) (ref.Named, error) {
	vars := mux.Vars(r)
	image := vars["image"]
	repo := vars["repo"]
	if repo == "" {
		repo = "library"
	}
	name := repo + "/" + image
	n, err := ref.WithName(name)
	if err != nil {
		return nil, fmt.Errorf("error parsing repository name: %v", err)
	}
	return n, nil
}
