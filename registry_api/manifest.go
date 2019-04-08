package api

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

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

	// get repository name
	var repositoryName ref.Named
	{
		image := vars["image"]
		repo := vars["repo"]
		if repo == "" {
			repo = "library"
		}
		name := repo + "/" + image
		n, err := ref.WithName(name)
		if err != nil {
			log.Printf("error parsing repository name: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		repositoryName = n
	}

	// get image reference
	var imageReference string
	var d digest.Digest
	{
		imageReference = vars["reference"]

		dgst, err := digest.Parse(imageReference)
		if err != nil {
			// We just have a tag
			//manifestHandler.Tag = reference
		} else {
			d = dgst
			//manifestHandler.Digest = dgst
		}

		_, err = ref.ParseAnyReference(imageReference)
		if err != nil {
			log.Printf("error parsing image reference: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// get manifest(-list) from backend registry
	serverBase := "https://registry-1.docker.io"
	transportAuth := makeHubTransport(serverBase, repositoryName.Name())
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
	tag := distribution.WithTag(imageReference)
	manifest, err := manifestService.Get(ctx, digest.Digest(""), tag)
	if err != nil {
		log.Printf("error getting manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	manifestList, isManifestList := manifest.(*manifestlist.DeserializedManifestList)
	if !isManifestList {
		log.Printf("unexpected menifest type: %v", reflect.TypeOf(manifest))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, m := range manifestList.Manifests {
		//p := m.Platform.
		_ = m
		//log.Printf("p: %v  urls: %v", p, m.URLs)
	}

	ct, p, err := manifest.Payload()
	if err != nil {
		return
	}

	log.Printf("Content-Type=%v", ct)

	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Length", fmt.Sprint(len(p)))
	w.Header().Set("Docker-Content-Digest", d.String())
	w.Header().Set("Etag", fmt.Sprintf(`"%s"`, d))
	w.Write(p)

	/*for _, d := range manifest.References() {
		log.Printf("descriptor: %v", d.Descriptor())
		//d.Descriptor()
	}*/

	// https://github.com/docker/distribution/tree/master/registry/client
	// https://github.com/docker/distribution/blob/master/registry/handlers/manifests.go
}
