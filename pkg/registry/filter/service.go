package filter

import (
	"context"
	"fmt"
	"io"
	"log"
	"metahub/pkg/registry"
	"metahub/pkg/storage"

	"github.com/docker/distribution"
	"github.com/opencontainers/go-digest"

	manifestListSchema "github.com/docker/distribution/manifest/manifestlist"
	manifestSchema "github.com/docker/distribution/manifest/schema2"
)

// https://docs.docker.com/registry/spec/manifest-v2-2/#image-manifest-field-descriptions
// https://docs.docker.com/registry/spec/api/#digest-parameter

func init() {
	_ = manifestListSchema.SchemaVersion
	_ = manifestSchema.SchemaVersion
}

type service struct {
	inner       registry.Service
	machineType storage.MachineType
}

// NewService returns a new registry service, wraping another registry Service, to filter image manifests
func NewService(inner registry.Service, machineType storage.MachineType) registry.Service {
	return &service{
		inner:       inner,
		machineType: machineType,
	}
}

func (s *service) GetBlob(ctx context.Context, repositoryString string, d digest.Digest) (io.ReadCloser, registry.Blob, error) {
	return s.inner.GetBlob(ctx, repositoryString, d)
}

func (s *service) GetManifest(ctx context.Context, repositoryString string, referenceString string) (registry.Manifest, error) {
	m, err := s.inner.GetManifest(ctx, repositoryString, referenceString)
	if err != nil {
		return m, err
	}
	unmarshaledManifest, _, err := distribution.UnmarshalManifest(m.ContentType, m.Data)
	if err != nil {
		return m, fmt.Errorf("error unmarshaling manifest: %v", err)
	}
	manifestList, ok := unmarshaledManifest.(*manifestListSchema.DeserializedManifestList)
	if !ok {
		return m, nil
	}
	manifestList, err = s.filterManifestsFromList(manifestList)
	if err != nil {
		return m, fmt.Errorf("error filtering manifest list: %v", err)
	}
	mediaType, payload, err := manifestList.Payload()
	if err != nil {
		return m, fmt.Errorf("error getting manifest payload: %v", err)
	}
	m = registry.Manifest{
		Data:        payload,
		ContentType: mediaType,
	}
	return m, nil
}

func (s *service) filterManifestsFromList(manifestList *manifestListSchema.DeserializedManifestList) (*manifestListSchema.DeserializedManifestList, error) {
	machineType := s.machineType
	log.Printf("machine type features: %v", machineType.Features)

	machineFeatureSet := make(map[string]struct{}, 0)
	for _, f := range machineType.Features {
		machineFeatureSet[f] = struct{}{}
	}

	filteredManifests := make([]manifestListSchema.ManifestDescriptor, 0)
	skipped := 0
	for _, m := range manifestList.Manifests {
		if len(m.Platform.Features) != len(machineFeatureSet) {
			skipped++
			//log.Printf("skipping manifest features %v", m.Platform.Features)
			continue
		}
		featureMismatch := false
		for i := 0; i < len(machineFeatureSet); i++ {
			platformFeature := m.Platform.Features[i]
			if _, ok := machineFeatureSet[platformFeature]; !ok {
				featureMismatch = true
				break
			}
		}
		if featureMismatch {
			skipped++
			log.Printf("skipping manifest features %v", m.Platform.Features)
			continue
		}
		//log.Printf("allow manifest features %v", m.Platform.Features)
		filteredManifests = append(filteredManifests, m)
	}
	log.Printf("skipped %d and left %d manifest features", skipped, len(filteredManifests))
	newManifestList, err := manifestListSchema.FromDescriptors(filteredManifests)
	if err != nil {
		return nil, fmt.Errorf("error generating new manifest list: %v", err)
	}
	return newManifestList, nil
}
