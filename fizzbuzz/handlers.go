package main

import (
	"fmt"
	"net/http"
	"errors"
	"strconv"
)

// Struct declaration
type FizzbuzzParams struct {
	Int1 int
	Int2 int
	Limit int
	Str1 string
	Str2 string
}


//Index
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint available:\nFizzbuzz generator: /fizzbuzz\n")
}

//FizzbuzzHandler functions

func isConformed(r *http.Request) (q FizzbuzzParams, err error){
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
			q.Str1 = r.FormValue("str1")
		}
		case "str2":
			q.Str2 = r.FormValue("str2")
		default: {
			err := errors.New("Bad parameter used\nPlease use these 5 keys: int1, int2, limit, str1, str2\nExample: http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz\n\n")
			return q, err
		}
		}
	}	
	// Return error if there are not 5 parameters or if the same key is used many times
	if nbValues == 0 {
		err := errors.New("To use Fizzbuzz, please add parameters to URL like this : http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz")
		return q, err
	} else if nbValues != 5 {
		// Nb of values incorrect or key used many times
		err := errors.New("Syntax error\nPlease use these 5 keys: int1, int2, limit, str1, str2\nExample: http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz\n")
		return q, err
	} 
	return q, err
}

func doFizzbuzz(q FizzbuzzParams) (fizzbuzzList string){
	for i := 1; i <= q.Limit; i++ {
		if i != 1 {
			fizzbuzzList = fizzbuzzList + ","
		} else if i == 1 {
			fizzbuzzList =  fizzbuzzList + "\""
		}
		switch {
		case i%q.Int1 == 0 && i%q.Int2 == 0: {

			fizzbuzzList = fizzbuzzList + q.Str1 + q.Str2
		}
		case i%q.Int1 == 0:
			fizzbuzzList = fizzbuzzList + q.Str1
		case i%q.Int2 == 0:
			fizzbuzzList = fizzbuzzList + q.Str2
		default:
			fizzbuzzList = fizzbuzzList + strconv.Itoa(i)
		}
		if i == q.Limit {
			fizzbuzzList = fizzbuzzList + "\".\n"
		}
	}
	return fizzbuzzList
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	// Check the conformity of the request
	q, err := isConformed(r)
	if err != nil{
		//fmt.Print(err) LOGS
		fmt.Fprintf(w, "Syntax Error\nPlease, use numbers for int1, int2 and limit values.\nError details below:\n%s", err)
		return
	}

	// Fizzbuzzlist generation
	fizzbuzzList := doFizzbuzz(q)

	// Print parameters defined in the URL 
	fmt.Fprintf(w, "Parameters defined:\n Int1: %d\n Int2: %d\n Limit: %d\n Str1: %s\n Str2: %s\n\n", q.Int1, q.Int2, q.Limit, q.Str1, q.Str2)
	
	// Print the fizzbuzzlist
	fmt.Fprintf(w, "Result:\n%s\n", fizzbuzzList)
}
