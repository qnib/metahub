package machinetypes

import (
	"encoding/json"
	"log"
	"metahub"
	"net/http"

	"github.com/gorilla/context"
)

func getDeleteHandler(env metahub.Environment) http.Handler {
	storageService := env.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		machineTypeService, err := storageService.MachineTypeService(ctx)
		if err != nil {
			log.Printf("failed to create MachineTypeService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := machineTypeService.Delete(accountName, requestParams.ID); err != nil {
			log.Printf("failed deleting machine type: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
