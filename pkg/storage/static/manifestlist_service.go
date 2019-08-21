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
func (s *manifestListService) Add(accountName string, ml *storage.ManifestList) (err error) {
	log.Println("Static backend does not support Add()")
	return
}

// Delete would normaly remove a new ManifestList - but in static that is not possible..
func (s *manifestListService) Delete(accountName string, id int64) (err error) {
	log.Println("Static backend does not support delete")
	return
}

// Update would normaly overwrite a new ManifestList - but in static that is not possible..
func (s *manifestListService) Update(accountName string, ml storage.ManifestList) (err error) {
	log.Println("Static backend does not support updates")
	return
}

// List is the only function that does something... :)
func (s *manifestListService) List(accountName string) ([]storage.ManifestList, error) {
	return getDummyManifestLists(), nil
}

// Get A manifestList from the local storage. nil if none is found
func (s *manifestListService) Get(accountName, repoName, tagName string) (ml *storage.ManifestList, err error) {
	for _, item := range getDummyManifestLists() {
		if item.RepoName == repoName && item.TagName == tagName {
			return &item, nil
		}
	}
	return nil, nil
}
