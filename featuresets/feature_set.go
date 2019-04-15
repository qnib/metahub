package featuresets

var featureSetEntityKind = "feature_set"

type featureSet struct {
	DisplayName string   `datastore:"name,noindex"`
	Features    []string `datastore:"features,noindex"`
	Login       string   `datastore:"login"`
	Password    string   `datastore:"password,noindex"`
}
