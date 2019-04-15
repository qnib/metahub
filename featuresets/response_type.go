package featuresets

type responseFeatureSet struct {
	ID          int64    `json:"id"`
	DisplayName string   `json:"name"`
	Features    []string `json:"features,omitempty"`
	Login       string   `json:"login"`
	Password    string   `json:"password"`
}
