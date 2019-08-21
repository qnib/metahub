package storage

import (
	"bytes"
	"encoding/json"
	"fmt"

	v1 "github.com/qnib/image-spec/specs-go/v2"
)

// repoName represents the full path without the registry component.
// https://docs.docker.com/registry/spec/api/
/******
Classically, repository names have always been two path components where each path component
is less than 30 characters. The V2 registry API does not enforce this.
The rules for a repository name are as follows:
1. A repository name is broken up into path components. A component of a
repository name must be at least one lowercase, alpha-numeric characters, optionally separated by periods,
dashes or underscores. More strictly, it must match the regular expression [a-z0-9]+(?:[._-][a-z0-9]+)*.
2. If a repository name has two or more path components, they must be separated by a forward slash (“/”).
3. The total length of a repository name, including slashes, must be less than 256 characters.
*/

// ManifestList describes the metahub-internal definitions of a ManifestList
type ManifestList struct {
	ID        int64      `json:"id"`
	RepoName  string     `json:"repo"`
	TagName   string     `json:"tag"`
	Manifests []Manifest `json:"manifests,omitempty"`
}

// ToBytes returns a bytearray
func (mtm ManifestList) ToBytes() []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(mtm)
	return reqBodyBytes.Bytes()
}

// NewManifestList generates a new Manifest list
func NewManifestList(repo, tag string, mfs ...Manifest) (ml ManifestList, err error) {
	ml.RepoName = repo
	ml.TagName = tag
	ml.Manifests = mfs
	return
}

// Manifest describes a single Image (aka Manifest)
type Manifest struct {
	ID       int64       `json:"id"`
	RepoName string      `json:"image"`
	TagName  string      `json:"tag"`
	Platform v1.Platform `json:"features,omitempty"`
}

// ToBytes returns a bytearray
func (mtm Manifest) ToBytes() []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(mtm)
	return reqBodyBytes.Bytes()
}

// NewManifest generates a new Manifest objects checking for correctness
func NewManifest(repo, tag, os, arch string, features ...string) (mf Manifest, err error) {
	mf.RepoName = repo
	mf.TagName = tag
	if arch != "amd64" && os != "linux" {
		return mf, fmt.Errorf("Let's first concentrate on 'linux' and 'amd64' - keep it simple")
	}
	mf.Platform = v1.Platform{
		Architecture: arch,
		OS:           os,
		Features:     features,
	}
	return
}
