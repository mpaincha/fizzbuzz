package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", index)
	router.HandleFunc("/fizzbuzz/", fizzbuzzHandler)
	router.HandleFunc("/fizzbuzz/statistics", fizzbuzzStatisticsHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}



// example of URL expexted : http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz
