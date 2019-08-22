package storage

import (
	"context"
	"io"
	"log"
	"metahub/pkg/registry"
	"metahub/pkg/storage"

	"github.com/opencontainers/go-digest"
)

type service struct {
	inner               registry.Service
	manifestListStorage storage.ManifestListService
}

// NewService returns a new registry service, wraping another registry Service, to filter image manifests
func NewService(inner registry.Service, mlService storage.ManifestListService) registry.Service {
	return &service{
		inner:               inner,
		manifestListStorage: mlService,
	}
}

func (s *service) GetBlob(ctx context.Context, repositoryString string, d digest.Digest) (io.ReadCloser, registry.Blob, error) {
	return s.inner.GetBlob(ctx, repositoryString, d)
}
func (s *service) GetManifest(ctx context.Context, repositoryString string, referenceString string) (m registry.Manifest, err error) {
	log.Printf("mlService: %v\n", s.manifestListStorage)
	ml, err := s.manifestListStorage.Get("qnib", repositoryString, referenceString)
	if err != nil {
		log.Printf("Got error while searching for ML in dynamic Storage: %s\n", err.Error())
		return m, err
	}
	if ml == nil {
		log.Printf("Could not find ManifestList in dynamic Storage; let's ask the next backend (dockerhub)\n")
		return s.inner.GetManifest(ctx, repositoryString, referenceString)
	}
	log.Printf("Got a ManifestList back (%v); but how to I createa  Manifest out of it?\n", ml)
	return
}
