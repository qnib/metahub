package proxy

import (
	"io"
	"log"
	"metahub/pkg/daemon"
	"net/http"
	"strconv"

	registryClient "github.com/docker/distribution/registry/client"
	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
)

func getBlobsHandler(service daemon.Service) http.Handler {
	//storageService := env.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
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

		// get backend blob service
		transportAuth := backendAuthTransport(serverBase, repositoryName.Name())
		repository, err := registryClient.NewRepository(repositoryName, serverBase, transportAuth)
		if err != nil {
			log.Printf("error creating repository object: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		blobService := repository.Blobs(ctx)

		// get blob stats
		blobStats, err := blobService.Stat(ctx, digest)
		if err != nil {
			log.Printf("error loading blob stats from backend: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// set response headers
		w.Header().Set("Content-Length", strconv.FormatInt(blobStats.Size, 10))
		w.Header().Set("Content-Type", blobStats.MediaType)
		w.Header().Set("Docker-Content-Digest", digest.String())
		w.Header().Set("Etag", digest.String())

		// open blob content stream
		blobContentReader, err := blobService.Open(ctx, digest)
		if err != nil {
			log.Printf("error getting blob stream from backend: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer blobContentReader.Close()

		// stream blob content to client
		_, err = io.CopyN(w, blobContentReader, blobStats.Size)
		if err != nil {
			log.Printf("error getting blob stream from backend: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
