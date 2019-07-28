package static

import (
	"context"
	"log"

	"metahub/pkg/storage"
)

type manifestListService struct {
	ctx context.Context
}

// Add would normaly add a new ManifestList - but in static that is not possible..
func (s *manifestListService) Add(ml *storage.ManifestList) (err error) {
	log.Println("Static backend does not support Add()")
	return
}

// Delete would normaly remove a new ManifestList - but in static that is not possible..
func (s *manifestListService) Delete(repoName, tagName string, id int64) (err error) {
	log.Println("Static backend does not support delete")
	return
}

// Update would normaly overwrite a new ManifestList - but in static that is not possible..
func (s *manifestListService) Update(ml *storage.ManifestList) (err error) {
	log.Println("Static backend does not support updates")
	return
}

// List is the only function that does something... :)
func (s *manifestListService) List() ([]*storage.ManifestList, error) {
	result := []*storage.ManifestList{}
	return []*storage.ManifestList{
		&mlQBench,
	}, nil
}
