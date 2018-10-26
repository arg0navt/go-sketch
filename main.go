package main

import (
	"log"
	"net/http"

	"./gosketch"
	"github.com/gorilla/mux"
)

func main() {
	s, err := gosketch.Read("../sketches/the-do-website-hero-templates.sketch")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.GetCSS)
	log.Fatal(http.ListenAndServe(":8080", router))
}
