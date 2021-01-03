package machinetypes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"
	"github.com/qnib/metahub/pkg/storage"

	"github.com/gorilla/context"
)

func getUpdateHandler(service daemon.Service) http.Handler {
	storageService := service.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountName := context.Get(r, "accountName").(string)

		decoder := json.NewDecoder(r.Body)
		var mt storage.MachineType
		err := decoder.Decode(&mt)
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

		if err := machineTypeService.Update(accountName, mt); err != nil {
			log.Printf("failed updating machine type: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		d, err := json.Marshal(mt)
		if err != nil {
			log.Printf("error marshaling response data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(d)
	})
}
