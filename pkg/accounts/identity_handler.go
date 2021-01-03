package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qnib/metahub/pkg/daemon"
	"github.com/qnib/metahub/pkg/storage"

	"github.com/gorilla/context"
)

func getIdentityHandler(service daemon.Service) http.Handler {
	authMiddleware := AuthMiddleware(service)

	return authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ctx := r.Context()
		accountName := context.Get(r, "accountName").(string)
		account := context.Get(r, "account").(*storage.Account)

		info := struct {
			AccountName string `json:"accountName"`
			DisplayName string `json:"displayName"`
		}{
			AccountName: accountName,
			DisplayName: account.DisplayName,
		}

		d, err := json.Marshal(info)
		if err != nil {
			log.Printf("error marshaling response data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(d)
	}))
}
