package main

import (
	"fmt"
	"net/http"
	"errors"
	"strconv"
	// "log"
	// "github.com/gorilla/mux"
)

// Struct declaration
type Query struct {
	Int1 int
	Int2 int
	Limit int
	Str1 string
	Str2 string
}


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to fizzbuzz")
}

func isConformed(r *http.Request) (q Query, err error){
	var nbValues int

	err = r.ParseForm()
	if err != nil {
		return q, err
	}

	// Check URL parameters
	for key := range r.Form {
		nbValues += 1

		switch key {
		case "int1": {
			q.Int1, err = strconv.Atoi(r.FormValue("int1"))
			if err != nil {
				return q, err
			}
		}
		case "int2": {
			q.Int2, err = strconv.Atoi(r.FormValue("int2"))
			if err != nil {
				return q, err
			}
		}
		case "limit": {
			q.Limit, err = strconv.Atoi(r.FormValue("limit"))
			if err != nil {
				return q, err
			}		
		}
		case "str1": {
			//check if is str
			q.Str1 = r.FormValue("str1")
		}
		case "str2":
			q.Str2 = r.FormValue("str2")
		default: {
			err := errors.New("bad paramters\n")
			return q, err
		}
		}
	}	
	// Return error if there are not 5 parameters or if the same key is used many times
	if nbValues != 5 {
		err := errors.New("nb of values incorrect or key used many times\n")
		fmt.Printf("nb value %d : %d %d %d %s %s\n", nbValues, q.Int1, q.Int2, q.Limit, q.Str1, q.Str2)
		return q, err
	}
	return q, err
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	// Check the conformity of the request
	q, err := isConformed(r)
	if err != nil{
		fmt.Print(err)
		fmt.Fprintf(w,"%s", err)
		return
	} else {
		fmt.Printf("Conformed\n")
	}
	fmt.Fprintf(w, "Parameters defined:\n Int1: %d\n Int2: %d\n Limit: %d\n Str1: %s\n Str2: %s\n", q.Int1, q.Int2, q.Limit, q.Str1, q.Str2)
}