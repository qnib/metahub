package storage

// ManifestListService provides an internal storage for ManifestLists defined within MetaHub
// to complement backend ManifestList or backends without ManifestList implementation
type ManifestListService interface {
	Add(accountName string, ml *ManifestList) (err error)
	Delete(accountName string, id int64) error
	List(accountName string) ([]ManifestList, error)
	Update(accountName string, ml ManifestList) error
	Get(accountName, repoName, tagName string) (ml *ManifestList, err error)
}
