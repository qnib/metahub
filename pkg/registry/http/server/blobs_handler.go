package server

import (
	"io"
	"log"
	"metahub/pkg/daemon"
	"metahub/pkg/registry"
	"net/http"
	"strconv"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
	"github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/redis"
	"time"
)

func getBlobsHandler(service daemon.Service) http.Handler {
	blobHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		// get repository and blob digest
		repository := getRepository(r)
		digest, err := digest.Parse(vars["reference"])
		if err != nil {
			log.Printf("error parsing blob reference: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// get manifest from registry service
		registryService := context.Get(r, "registryService").(registry.Service)
		blobReader, blob, err := registryService.GetBlob(ctx, repository, digest)
		if err != nil {
			log.Printf("error getting blob: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer blobReader.Close()

		// set response headers
		w.Header().Set("Content-Length", strconv.FormatInt(blob.Size, 10))
		w.Header().Set("Content-Type", blob.MediaType)
		w.Header().Set("Docker-Content-Digest", digest.String())
		w.Header().Set("Etag", digest.String())

		// stream blob content to client
		_, err = io.CopyN(w, blobReader, blob.Size)
		if err != nil {
			log.Printf("error getting blob stream from backend: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	return blobHandler
}


func getCachedBlobsHandler(service daemon.Service) http.Handler {
	blobHandler := getBlobsHandler(service)

	ringOpt := &redis.RingOptions{
		Addrs: map[string]string{
			"server": ":6379",
		},
	}
	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(redis.NewAdapter(ringOpt)),
		cache.ClientWithTTL(10 * time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		log.Printf("error getting cache client: %v", err)
		return blobHandler
	}
	log.Println("Create cached blobHandler")
	return cacheClient.Middleware(blobHandler)
}
