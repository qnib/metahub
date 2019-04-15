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
		Name string `json:"name"`
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
	featureSetKey := datastore.NameKey(featureSetEntityKind, requestParams.Name, accountKey)

	var fs featureSet
	fs.Features = []string{
		"gpu",
		"bla",
	}
	if _, err := datastoreClient.Put(ctx, featureSetKey, &fs); err != nil {
		log.Printf("error putting feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := struct {
		Name     string   `json:"name"`
		Features []string `json:"features"`
	}{
		Name:     requestParams.Name,
		Features: fs.Features,
	}
	d, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error marshaling feature set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
