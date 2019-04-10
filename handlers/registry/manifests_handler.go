package registry

import (
	"fmt"
	"log"
	"net/http"

	"github.com/docker/distribution"
	manifestlist "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema2 "github.com/docker/distribution/manifest/schema2"
	registryClient "github.com/docker/distribution/registry/client"
	"github.com/gorilla/mux"
	digestLib "github.com/opencontainers/go-digest"
)

// https://docs.docker.com/registry/spec/manifest-v2-2/#image-manifest-field-descriptions
// https://docs.docker.com/registry/spec/api/#digest-parameter

func init() {
	manifestlist.FromDescriptors([]manifestlist.ManifestDescriptor{})
	_ = manifestSchema2.SchemaVersion
}

func manifestsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	// get repository name
	repositoryName, err := getRepository(r)
	if err != nil {
		log.Printf("error parsing repository: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get image reference
	var tag distribution.ManifestServiceOption
	var digest digestLib.Digest
	{
		reference := vars["reference"]
		dgst, err := digestLib.Parse(reference)
		if err != nil {
			tag = distribution.WithTag(reference)
		} else {
			digest = dgst
		}
	}

	// get manifest(-list) from backend registry
	serverBase := "https://registry-1.docker.io"
	transportAuth := makeHubTransport(serverBase, repositoryName.Name())
	repository, err := registryClient.NewRepository(repositoryName, serverBase, transportAuth)
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
	var manifest distribution.Manifest
	if tag == nil {
		manifest, err = manifestService.Get(ctx, digest)
	} else {
		manifest, err = manifestService.Get(ctx, digest, tag)
	}
	if err != nil {
		log.Printf("error getting backend manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	manifestList, isManifestList := manifest.(*manifestlist.DeserializedManifestList)
	if isManifestList {
		for i, m := range manifestList.Manifests {
			if m.Platform.OS != "windows" || m.Platform.Architecture != "amd64" {
				continue
			}
			manifestList.Manifests[i] = m
		}
		newManifestList, err := manifestlist.FromDescriptors(manifestList.Manifests)
		if err != nil {
			log.Printf("error generating new manifest list: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		manifest = newManifestList
	}

	_, isManifest := manifest.(*manifestSchema2.DeserializedManifest)
	if isManifest {
	}

	ct, p, err := manifest.Payload()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Length", fmt.Sprint(len(p)))
	w.Header().Set("Docker-Content-Digest", digest.String())
	w.Header().Set("Etag", fmt.Sprintf(`"%s"`, digest))
	w.Write(p)

	/*for _, d := range manifest.References() {
		log.Printf("descriptor: %v", d.Descriptor())
		//d.Descriptor()
	}*/

	// https://github.com/docker/distribution/tree/master/registry/client
	// https://github.com/docker/distribution/blob/master/registry/handlers/manifests.go
}
