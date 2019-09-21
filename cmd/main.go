package main

import (
	// "fmt"
	// "html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/", index)
	router.HandleFunc("/fizzbuzz/", fizzbuzzHandler)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}



// example of URL expexted : http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz

// Default value
/*const (
	int1 = 3
	int2 = 5
	limit = 100
	str1 = "fizz"
	str2 = "buzz"	
)*/
