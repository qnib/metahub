package featuresets

import (
	"encoding/json"
	"log"
	"metahub/auth"
	"net/http"

	"cloud.google.com/go/datastore"

	"github.com/gorilla/context"
)

func add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountName := context.Get(r, "account").(string)

	decoder := json.NewDecoder(r.Body)
	var requestParams struct {
		DisplayName string   `json:"name"`
		Features    []string `json:"features"`
	}
	err := decoder.Decode(&requestParams)
	if err != nil {
		log.Printf("error decoding request data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	datastoreClient, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Printf("failed to create client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accountKey := datastore.NameKey(auth.AccountEntityKind, accountName, nil)
	featureSetKey := datastore.IncompleteKey(featureSetEntityKind, accountKey)

	fs := featureSet{
		DisplayName: requestParams.DisplayName,
		Features:    requestParams.Features,
		Login:       "test test test",
		Password:    "password",
	}
	featureSetKey, err = datastoreClient.Put(ctx, featureSetKey, &fs)
	if err != nil {
		log.Printf("error putting feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("featureSetKey: %v", featureSetKey)

	responseData := responseFeatureSet{
		ID:          featureSetKey.ID,
		DisplayName: fs.DisplayName,
		Features:    fs.Features,
		Login:       fs.Login,
		Password:    fs.Password,
	}
	d, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error marshaling feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
