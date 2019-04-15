package featuresets

import (
	"encoding/json"
	"log"
	"metahub/auth"
	"net/http"

	"cloud.google.com/go/datastore"

	"github.com/gorilla/context"
)

func list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountName := context.Get(r, "account").(string)

	datastoreClient, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Printf("failed to create client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accountKey := datastore.NameKey(auth.AccountEntityKind, accountName, nil)
	var featureSets []featureSet
	q := datastore.NewQuery(featureSetEntityKind)
	q = q.Ancestor(accountKey)
	featureSetKeys, err := datastoreClient.GetAll(ctx, q, &featureSets)
	if err != nil {
		log.Printf("error querying feature sets: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("%d feature sets", len(featureSets))

	type responseFeatureSet struct {
		Name     string   `json:"name"`
		Features []string `json:"features,omitempty"`
	}

	var responseData struct {
		FeatureSets []responseFeatureSet `json:"featureSets,omitempty"`
	}
	for i, fs := range featureSets {
		responseData.FeatureSets = append(responseData.FeatureSets, responseFeatureSet{
			Name:     featureSetKeys[i].Name,
			Features: fs.Features,
		})
	}

	d, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error marshaling feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
