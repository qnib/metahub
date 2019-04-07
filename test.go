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

    //Accept header
    // application/vnd.docker.distribution.manifest.v2+json
    // application/vnd.docker.distribution.manifest.list.v2+json
    // return Content-Type

    // https://docs.docker.com/registry/spec/api/#manifest

    // /v2/test/manifests/latest
}
