package api

import (
	"log"
	"net/http"
	"time"

	"github.com/docker/distribution"

	ref "github.com/docker/distribution/reference"
	client "github.com/docker/distribution/registry/client"
	"github.com/docker/distribution/registry/client/auth"
	"github.com/docker/distribution/registry/client/auth/challenge"
	"github.com/docker/distribution/registry/client/transport"
	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
)

func makeHubTransport(server, image string) http.RoundTripper {
	base := http.DefaultTransport

	modifiers := []transport.RequestModifier{
		transport.NewHeaderRequestModifier(http.Header{
			"User-Agent": []string{"my-client"},
		}),
	}

	authTransport := transport.NewTransport(base, modifiers...)
	pingClient := &http.Client{
		Transport: authTransport,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", server+"/v2/", nil)
	if err != nil {
		panic(err)
	}

	challengeManager := challenge.NewSimpleManager()
	resp, err := pingClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := challengeManager.AddResponse(resp); err != nil {
		panic(err)
	}
	tokenHandler := auth.NewTokenHandler(base, nil, image, "pull")
	modifiers = append(modifiers, auth.NewAuthorizer(challengeManager, tokenHandler, auth.NewBasicHandler(nil)))

	return transport.NewTransport(base, modifiers...)
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

	// "Accept" header
	// application/vnd.docker.distribution.manifest.v2+json
	// application/vnd.docker.distribution.manifest.list.v2+json
	// return "Content-Type" header
	// https://docs.docker.com/registry/spec/api/#manifest
	//log.Printf("name:%s  reference:%s", name, reference)

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

	//distribution.ManifestMediaTypes()

	manifest, err := manifestService.Get(ctx, digest.Digest(""), distribution.WithTag(reference)) /*distribution.WithManifestMediaTypes([]string{
		"application/vnd.docker.distribution.manifest.v2+json",
		"application/vnd.docker.distribution.manifest.list.v2+json",
	})*/
	if err != nil {
		log.Printf("error getting manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("manifests: %v", manifest)

	// https://github.com/docker/distribution/tree/master/registry/client
	// https://github.com/docker/distribution/blob/master/registry/handlers/manifests.go
}
