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
	inner      registry.Service
	storageSvc storage.Service
}

// NewService returns a new registry service, wraping another registry Service, to filter image manifests
func NewService(inner registry.Service, storageSvc storage.Service) registry.Service {
	return &service{
		inner:      inner,
		storageSvc: storageSvc,
	}
}

func (s *service) GetBlob(ctx context.Context, repositoryString string, d digest.Digest) (io.ReadCloser, registry.Blob, error) {
	return s.inner.GetBlob(ctx, repositoryString, d)
}
func (s *service) GetManifest(ctx context.Context, repositoryString string, referenceString string) (m registry.Manifest, err error) {
	/*
		// Checks if the incoming request can be served with a dynamic ML
		ml, err := s.storageSvc.GetManifest("qnib", repositoryString, referenceString)
		if err != nil {
			return m, err
		}
		if ml == nil {
			// If no dynamic ML is found, ask the inner registry
			return s.inner.GetManifest(ctx, repositoryString, referenceString)
		}
		log.Printf("Found ManifestList locally: %v\n", ml)
	*/
	log.Printf("Skip Dynamic Layer for: %s:%s\n", repositoryString, referenceString)
	return s.inner.GetManifest(ctx, repositoryString, referenceString)
}
