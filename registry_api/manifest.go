package api

import (
	"log"
	"net/http"

	"github.com/docker/distribution"

	manifestlist "github.com/docker/distribution/manifest/manifestlist"
	manifest "github.com/docker/distribution/manifest/schema2"
	ref "github.com/docker/distribution/reference"
	client "github.com/docker/distribution/registry/client"
	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
)

//distribution.RegisterManifestSchema

func init() {
	manifestlist.FromDescriptors([]manifestlist.ManifestDescriptor{})
	_ = manifest.SchemaVersion
}

func manifestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	image := vars["image"]
	repo := vars["repo"]
	if repo == "" {
		repo = "library"
	}
	name := repo + "/" + image
	reference := vars["reference"]

	//repositoryName, err := ref.WithName("index.docker.io/" + name)
	repositoryName, err := ref.WithName(name)
	if err != nil {
		log.Printf("error parsing repo name: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serverBase := "https://registry-1.docker.io"

	transportAuth := makeHubTransport(serverBase, name)

	repository, err := client.NewRepository(repositoryName, serverBase, transportAuth)
	if err != nil {
		log.Printf("error creating repository object: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	manifestService, err := repository.Manifests(ctx)
	if err != nil {
		log.Printf("error creating repository object: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	manifest, err := manifestService.Get(ctx, digest.Digest(""), distribution.WithTag(reference))
	if err != nil {
		log.Printf("error getting manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, d := range manifest.References() {
		log.Printf("descriptor: %v", d.Descriptor())
	}

	// https://github.com/docker/distribution/tree/master/registry/client
	// https://github.com/docker/distribution/blob/master/registry/handlers/manifests.go
}
