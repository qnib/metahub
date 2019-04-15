package featuresets

import (
	"encoding/json"
	"log"
	"metahub/auth"
	"net/http"

	"cloud.google.com/go/datastore"

	"github.com/gorilla/context"
)

func delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountName := context.Get(r, "account").(string)

	decoder := json.NewDecoder(r.Body)
	var requestParams struct {
		ID int64 `json:"id"`
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
	featureSetKey := datastore.IDKey(featureSetEntityKind, requestParams.ID, accountKey)

	log.Printf("featureSetKey: %v", featureSetKey)

	err = datastoreClient.Delete(ctx, featureSetKey)
	if err != nil {
		log.Printf("error deleting feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
