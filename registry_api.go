package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/url"

	"github.com/docker/distribution"

	ref "github.com/docker/distribution/reference"
	client "github.com/docker/distribution/registry/client"
	_ "github.com/docker/distribution/registry/storage/cache"
	_ "github.com/docker/distribution/registry/storage/cache/memory"
	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/v2", baseHandler).Methods("GET")
	r.HandleFunc("/v2/{name}/manifests/{reference}", manifestHandler).Methods("GET")
	http.Handle("/v2/", r)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
}

type cred struct {
}

func (c *cred) Basic(*url.URL) (string, string) {
	log.Printf("baaasic")
	return "hovu96", "test"
}

func (c *cred) RefreshToken(*url.URL, string) string {
	return ""
}

func (c *cred) SetRefreshToken(realm *url.URL, service, token string) {

}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func manifestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	name := vars["name"]
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

	var transportAuth http.RoundTripper
	//challengeManager := challenge.NewSimpleManager()
	//authenticationHandler := auth.NewBasicHandler(&cred{})
	//authorizer := auth.NewAuthorizer(challengeManager, authenticationHandler)
	//authHeader := http.Header{}
	//authHeader.Set("Authorization", "Basic "+basicAuth("user", "pass"))
	//transportAuth := transport.NewTransport(nil, transport.NewHeaderRequestModifier(authHeader))

	repository, err := client.NewRepository(repositoryName, "https://registry-1.docker.io/", transportAuth)
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

	manifest, err := manifestService.Get(ctx, digest.Digest(""), distribution.WithTag(reference), distribution.WithManifestMediaTypes([]string{
		//"application/vnd.docker.distribution.manifest.v2+json",
		"application/vnd.docker.distribution.manifest.list.v2+json",
	}))
	if err != nil {
		log.Printf("error getting manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("manifests: %v", manifest)

	// https://github.com/docker/distribution/tree/master/registry/client
	// https://github.com/docker/distribution/blob/master/registry/handlers/manifests.go
}
