package registry

import (
	"context"
	"io"

	digest "github.com/opencontainers/go-digest"
)

// Service provides access to a remote registry image repositories
type Service interface {
	GetBlob(ctx context.Context, repository string, d digest.Digest) (io.ReadCloser, Blob, error)
	GetManifest(ctx context.Context, repositoryString string, referenceString string) (Manifest, error)
}
