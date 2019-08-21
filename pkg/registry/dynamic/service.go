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
	inner        registry.Service
	ManifestList storage.ManifestListService
}

// NewService returns a new registry service, wraping another registry Service, to filter image manifests
func NewService(inner registry.Service, manifestList storage.ManifestListService) registry.Service {
	return &service{
		inner:        inner,
		ManifestList: manifestList,
	}
}

func (s *service) GetBlob(ctx context.Context, repositoryString string, d digest.Digest) (io.ReadCloser, registry.Blob, error) {
	return s.inner.GetBlob(ctx, repositoryString, d)
}
func (s *service) GetManifest(ctx context.Context, repositoryString string, referenceString string) (m registry.Manifest, err error) {
	ml, err := s.ManifestList.Get("qnib", repositoryString, referenceString)
	if err != nil {
		return m, err
	}
	if ml == nil {
		return s.inner.GetManifest(ctx, repositoryString, referenceString)
	}
	log.Printf("Found ManifestList locally: %v\n", ml)
	return s.inner.GetManifest(ctx, repositoryString, referenceString)
}
