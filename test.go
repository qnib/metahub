package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleTest)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	fmt.Fprint(w, "test")
}
