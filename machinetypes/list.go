package machinetypes

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
	var machineTypes []machineType
	q := datastore.NewQuery(machineTypeEntityKind)
	q = q.Ancestor(accountKey)
	machineTypeKeys, err := datastoreClient.GetAll(ctx, q, &machineTypes)
	if _, ok := err.(*datastore.ErrFieldMismatch); ok {
		err = nil
	}
	if err != nil {
		log.Printf("error querying feature sets: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//log.Printf("%d feature sets", len(featureSets))

	var responseData struct {
		MachineTypes []responseMachineType `json:"machineTypes,omitempty"`
	}
	for i, fs := range machineTypes {
		k := machineTypeKeys[i]
		responseData.MachineTypes = append(responseData.MachineTypes, responseMachineType{
			ID:          k.ID,
			DisplayName: fs.DisplayName,
			Features:    fs.Features,
			Login:       k.Encode(),
			Password:    fs.Password,
		})
	}

	d, err := json.Marshal(responseData)
	if err != nil {
		log.Printf("error marshaling response data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
