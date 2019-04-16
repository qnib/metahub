package machinetypes

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
	machineTypeKey := datastore.IncompleteKey(machineTypeEntityKind, accountKey)

	newPassword, err := generateLoginPassword()
	if err != nil {
		log.Printf("failed to generate new password: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mt := machineType{
		DisplayName: requestParams.DisplayName,
		Features:    requestParams.Features,
		Password:    newPassword,
	}
	machineTypeKey, err = datastoreClient.Put(ctx, machineTypeKey, &mt)
	if err != nil {
		log.Printf("error putting machine type: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("machineTypeKey: %v", machineTypeKey)

	responseData := responseMachineType{
		ID:          machineTypeKey.ID,
		DisplayName: mt.DisplayName,
		Features:    mt.Features,
		Login:       machineTypeKey.String(),
		Password:    mt.Password,
	}
	d, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error marshaling response data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
