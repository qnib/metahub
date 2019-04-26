package machinetypes

import (
	"encoding/json"
	"log"
	"metahub/pkg/daemon"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
)

func getGetHandler(service daemon.Service) http.Handler {
	storageService := service.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountName := context.Get(r, "account").(string)

		machineTypeIDString := r.URL.Query().Get("id")
		machineTypeID, _ := strconv.ParseInt(machineTypeIDString, 10, 64)

		machineTypeService, err := storageService.MachineTypeService(ctx)
		if err != nil {
			log.Printf("failed to create MachineTypeService: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mt, err := machineTypeService.GetByID(accountName, machineTypeID)
		if err != nil {
			log.Printf("failed getting machine type: %v", err)
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
