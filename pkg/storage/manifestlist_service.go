package storage

// ManifestListService provides an internal storage for ManifestLists defined within MetaHub
// to complement backend ManifestList or backends without ManifestList implementation
type ManifestListService interface {
	HasManifestList(repoName string) (result bool, err error)
	Add(ml *ManifestList) (err error)
	Delete(repoName, tagName string, id int64) error
	List() ([]*ManifestList, error)
	Update(ml *ManifestList) error
}
