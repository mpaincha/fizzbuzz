package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	// Router initialization
	router := mux.NewRouter().StrictSlash(true)

	// Registration of two routes mapping URL paths to handlers
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/fizzbuzz/", fizzbuzzHandler)

	// Port used 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}