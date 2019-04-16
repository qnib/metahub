package machinetypes

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
	machineTypeKey := datastore.IDKey(machineTypeEntityKind, requestParams.ID, accountKey)

	log.Printf("machineTypeKey: %v", machineTypeKey)

	err = datastoreClient.Delete(ctx, machineTypeKey)
	if err != nil {
		log.Printf("error deleting feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
