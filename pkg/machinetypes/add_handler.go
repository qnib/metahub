package machinetypes

import (
	"encoding/json"
	"log"
	"metahub/pkg/daemon"
	"net/http"

	"metahub/pkg/storage"

	"github.com/gorilla/context"
)

func getAddHandler(service daemon.Service) http.Handler {
	storageService := service.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		machineTypeService, err := storageService.MachineTypeService(ctx)
		if err != nil {
			log.Printf("failed to create MachineTypeService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		newPassword, err := generateLoginPassword()
		if err != nil {
			log.Printf("failed to generate new password: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mt := storage.MachineType{
			DisplayName: requestParams.DisplayName,
			Features:    requestParams.Features,
			Password:    newPassword,
		}

		if err := machineTypeService.Add(accountName, &mt); err != nil {
			log.Printf("failed adding machine type: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		d, err := json.Marshal(mt)
		if err != nil {
			log.Printf("error marshaling response data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(d)
	})
}
