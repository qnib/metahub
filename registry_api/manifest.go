package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/docker/distribution"

	manifestlist "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema2 "github.com/docker/distribution/manifest/schema2"
	ref "github.com/docker/distribution/reference"
	client "github.com/docker/distribution/registry/client"
	"github.com/gorilla/mux"
	digest "github.com/opencontainers/go-digest"
)

//distribution.RegisterManifestSchema

func init() {
	manifestlist.FromDescriptors([]manifestlist.ManifestDescriptor{})
	_ = manifestSchema2.SchemaVersion
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
	//var iirrr ref.Reference
	var d digest.Digest
	var isDigest bool
	{
		imageReference = vars["reference"]

		dgst, err := digest.Parse(imageReference)
		if err != nil {
			// We just have a tag
			//manifestHandler.Tag = reference
		} else {
			d = dgst
			isDigest = true
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
	var manifest distribution.Manifest
	if isDigest {
		manifest, err = manifestService.Get(ctx, d)
	} else {
		manifest, err = manifestService.Get(ctx, digest.Digest(""), tag)
	}
	if err != nil {
		log.Printf("error getting manifest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mm, isManifest := manifest.(*manifestSchema2.DeserializedManifest)
	if isManifest {
		//log.Printf("sha: %v", mm.)
		for _, l := range mm.Layers {
			if len(l.URLs) == 0 {
				/*l.URLs = []string{
					"http://localhost:8080/test",
				}*/
			}
			//log.Printf("layer urls: %v sha: %v", l.URLs, l.Digest)
		}
		//url := fmt.Sprintf("%s%s", serverBase, r.URL.Path)
		//log.Printf("url: %s", url)
		//http.Redirect(w, r, url, http.StatusMovedPermanently)
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}

	manifestList, isManifestList := manifest.(*manifestlist.DeserializedManifestList)
	if isManifestList {
		// TODO: liste filtern
		/*for _, m := range manifestList.Manifests {
			p := m.Platform.
		}*/
		manifestList.References()
	}

	ct, p, err := manifest.Payload()
	if err != nil {
		return
	}

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
