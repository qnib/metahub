package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"
	"github.com/qnib/metahub/pkg/registry"

	"github.com/gorilla/context"

	manifestListSchema "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema "github.com/docker/distribution/manifest/schema2"
	"github.com/gorilla/mux"
	digestLib "github.com/opencontainers/go-digest"
)

// https://docs.docker.com/registry/spec/api/#digest-parameter

func init() {
	_ = manifestListSchema.SchemaVersion
	_ = manifestSchema.SchemaVersion
}

func getRegistryHandler(service daemon.Service) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		// get repository and image reference
		repository := getRepository(r)
		reference := vars["reference"]

		// get manifest from registry service
		registryService := context.Get(r, "registryService").(registry.Service)
		manifest, err := registryService.GetManifest(ctx, repository, reference)
		if err != nil {
			log.Printf("error getting manifest: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// send response
		w.Header().Set("Content-Type", manifest.ContentType)
		w.Header().Set("Content-Length", fmt.Sprint(len(manifest.Data)))
		digest, _ := digestLib.Parse(reference)
		w.Header().Set("Docker-Content-Digest", digest.String())
		w.Header().Set("Etag", fmt.Sprintf(`"%s"`, digest))
		w.Write(manifest.Data)
	})
}
