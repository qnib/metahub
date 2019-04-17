package machinetypes

import (
	"encoding/json"
	"log"
	"metahub"
	"metahub/pkg/storage"
	"net/http"

	"github.com/gorilla/context"
)

func getListHandler(env metahub.Environment) http.Handler {
	storageService := env.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountName := context.Get(r, "account").(string)

		machineTypeService, err := storageService.MachineTypeService(ctx)
		if err != nil {
			log.Printf("failed to create MachineTypeService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		machineTypes, err := machineTypeService.List(accountName)
		if err != nil {
			log.Printf("error querying machine types: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//log.Printf("%d feature sets", len(featureSets))

		response := struct {
			MachineTypes []storage.MachineType `json:"machineTypes,omitempty"`
		}{
			MachineTypes: machineTypes,
		}

		d, err := json.Marshal(response)
		if err != nil {
			log.Printf("error marshaling response data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(d)

	})
}
