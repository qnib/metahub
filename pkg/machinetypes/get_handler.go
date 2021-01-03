package machinetypes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/qnib/metahub/pkg/daemon"

	"github.com/gorilla/context"
)

func getGetHandler(service daemon.Service) http.Handler {
	storageService := service.Storage()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		accountName := context.Get(r, "accountName").(string)

		machineTypeIDString := r.URL.Query().Get("id")
		machineTypeID, err := strconv.ParseInt(machineTypeIDString, 10, 64)
		if err != nil {
			log.Printf("failed parsing machine type id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
