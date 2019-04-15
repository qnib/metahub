package featuresets

var featureSetEntityKind = "feature_set"

type featureSet struct {
	Features []string `datastore:"features,noindex"`
}
