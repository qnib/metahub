package auth

import (
	"io/ioutil"
	"log"
	"net/http"
)

func githubHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
