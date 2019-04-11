package registry

import (
	"fmt"
	"log"
	"net/http"

	"github.com/docker/distribution"
	manifestListSchema "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema "github.com/docker/distribution/manifest/schema2"
	registryClient "github.com/docker/distribution/registry/client"
	"github.com/gorilla/mux"
	digestLib "github.com/opencontainers/go-digest"
)

// https://docs.docker.com/registry/spec/manifest-v2-2/#image-manifest-field-descriptions
// https://docs.docker.com/registry/spec/api/#digest-parameter

func init() {
	_ = manifestListSchema.SchemaVersion
	_ = manifestSchema.SchemaVersion
}

func (reg registry) manifestsHandler(w http.ResponseWriter, r *http.Request) {
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

	// get manifest from backend registry
	transportAuth := backendAuthTransport(serverBase, repositoryName.Name())
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

	manifestList, isManifestList := manifest.(*manifestListSchema.DeserializedManifestList)
	if isManifestList {
		newManifestList, err := filterManifestsFromList(manifestList)
		if err != nil {
			log.Printf("error filtering manifest list: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		manifest = newManifestList
	}

	_, isManifest := manifest.(*manifestSchema.DeserializedManifest)
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
}

func filterManifestsFromList(menifestList *manifestListSchema.DeserializedManifestList) (*manifestListSchema.DeserializedManifestList, error) {
	for i, m := range menifestList.Manifests {
		if m.Platform.OS != "windows" || m.Platform.Architecture != "amd64" {
			continue
		}
		menifestList.Manifests[i] = m
	}
	newManifestList, err := manifestListSchema.FromDescriptors(menifestList.Manifests)
	if err != nil {
		return nil, fmt.Errorf("error generating new manifest list: %v", err)
	}
	return newManifestList, nil
}
