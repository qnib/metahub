package client

import (
	"context"
	"fmt"
	"io"

	"github.com/qnib/metahub/pkg/registry"

	"github.com/docker/distribution"
	"github.com/docker/distribution/reference"
	registryClient "github.com/docker/distribution/registry/client"
	"github.com/opencontainers/go-digest"

	manifestListSchema "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema "github.com/docker/distribution/manifest/schema2"
)

// https://docs.docker.com/registry/spec/manifest-v2-2/#image-manifest-field-descriptions

func init() {
	_ = manifestListSchema.SchemaVersion
	_ = manifestSchema.SchemaVersion
}

type service struct {
	serverBase string
}

// NewService returns a new HTTP client registry service
func NewService() registry.Service {
	return &service{
		serverBase: "https://registry-1.docker.io",
	}
}

func (s *service) newRepositoryClient(repositoryString string) (distribution.Repository, error) {

	//TODO: add "library" segment?
	repositoryName, err := reference.WithName(repositoryString)
	if err != nil {
		return nil, fmt.Errorf("error parsing repository name: %v", err)
	}

	// get backend blob service
	transportAuth := backendAuthTransport(serverBase, repositoryString)
	repositoryClient, err := registryClient.NewRepository(repositoryName, s.serverBase, transportAuth)
	if err != nil {
		return nil, fmt.Errorf("error creating repository object: %v", err)
	}

	return repositoryClient, nil
}

func (s *service) GetBlob(ctx context.Context, repositoryString string, d digest.Digest) (io.ReadCloser, registry.Blob, error) {
	var blob registry.Blob

	// get backend repository blobs service
	repositoryClient, err := s.newRepositoryClient(repositoryString)
	if err != nil {
		return nil, blob, fmt.Errorf("error loading blob stats from backend: %v", err)
	}
	blobService := repositoryClient.Blobs(ctx)

	// get blob stats
	blobStats, err := blobService.Stat(ctx, d)
	if err != nil {
		return nil, blob, fmt.Errorf("error loading blob stats from backend: %v", err)
	}

	blob.Size = blobStats.Size
	blob.MediaType = blobStats.MediaType

	// open blob content stream
	blobContentReader, err := blobService.Open(ctx, d)
	if err != nil {
		return nil, blob, fmt.Errorf("error getting blob stream from backend: %v", err)
	}

	return blobContentReader, blob, nil
}

func (s *service) GetManifest(ctx context.Context, repositoryString string, referenceString string) (registry.Manifest, error) {
	var m registry.Manifest

	// get image reference
	var tag distribution.ManifestServiceOption
	var d digest.Digest
	{
		dgst, err := digest.Parse(referenceString)
		if err != nil {
			tag = distribution.WithTag(referenceString)
		} else {
			d = dgst
		}
	}

	// get backend manifest service
	repositoryClient, err := s.newRepositoryClient(repositoryString)
	if err != nil {
		return m, fmt.Errorf("error loading blob stats from backend: %v", err)
	}
	manifestService, err := repositoryClient.Manifests(ctx)
	if err != nil {
		return m, fmt.Errorf("error creating repository object: %v", err)
	}

	// call backend manifest
	var manifest distribution.Manifest
	if tag == nil {
		manifest, err = manifestService.Get(ctx, d)
	} else {
		manifest, err = manifestService.Get(ctx, d, tag)
	}
	if err != nil {
		return m, fmt.Errorf("error getting backend manifest: %v", err)
	}

	mediaType, payload, err := manifest.Payload()
	if err != nil {
		return m, fmt.Errorf("error getting manifest payload: %v", err)
	}
	m = registry.Manifest{
		Data:        payload,
		ContentType: mediaType,
	}
	return m, nil
}
