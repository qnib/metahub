package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qnib/metahub/pkg/daemon"
	registry "github.com/qnib/metahub/pkg/registry/http/client"
	"github.com/qnib/metahub/pkg/storage/boltdb"

	"github.com/qnib/metahub/cmd"
)

var (
	version = flag.Bool("version", false, "print version")
	config  = flag.String("config", "", "Config holding the initial machine types and users")
	genHash = flag.String("genhash", "", "Create password hash and exit")
)

// Log intercepts each requests and writes it out
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if *version {
		fmt.Println(`v0.2.10`)
		os.Exit(0)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	storageService := boltdb.NewService(*config)
	if *genHash != "" {
		mtScv, _ := storageService.MachineTypeService(context.Background())
		fmt.Println(mtScv.GenPasswordHash(*genHash))
		os.Exit(0)
	}

	registryService := registry.NewService()
	daemonService := daemon.NewService(storageService, registryService)

	router := cmd.RegisterRoutes(daemonService)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), Log(router)))
}
