package main

import (
	"log"
	"net/http"

	"./gosketch"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", gosketch.Read)
	log.Fatal(http.ListenAndServe(":8080", router))
}
