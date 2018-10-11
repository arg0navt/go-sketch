package main

import (
	"log"
	"net/http"

	"./gosketch"
	"github.com/gorilla/mux"
)

func main() {
	i, err := gosketch.Read("./progressive-web-app-onboarding-richcullen.sketch")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", i.GetCSS)
	log.Fatal(http.ListenAndServe(":8080", router))
}
