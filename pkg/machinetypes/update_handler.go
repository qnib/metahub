package machinetypes

import (
	"encoding/json"
	"log"
	"metahub/pkg/daemon"
	"metahub/pkg/storage"
	"net/http"

	"github.com/gorilla/context"
)

func getUpdateHandler(env daemon.Environment) http.Handler {
	storageService := env.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountName := context.Get(r, "account").(string)

		decoder := json.NewDecoder(r.Body)
		var requestParams struct {
			ID          int64    `json:"id"`
			DisplayName string   `json:"name"`
			Features    []string `json:"features"`
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

		mt := storage.MachineType{
			ID:          requestParams.ID,
			DisplayName: requestParams.DisplayName,
			Features:    requestParams.Features,
		}

		if err := machineTypeService.Update(accountName, mt); err != nil {
			log.Printf("failed updating machine type: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
