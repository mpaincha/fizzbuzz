package main

import (
	"fmt"
	"net/http"
	"errors"
	"strconv"
)

// Declaration of the struct which will be used to store URL queries
type FizzbuzzParams struct {
	Int1 int
	Int2 int
	Limit int
	Str1 string
	Str2 string
}


// Handler "Index" which is used when there is no specific path. There is only the protocol and domain as http://localhost:8080/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint available:\nFizzbuzz generator: /fizzbuzz\n")
}

// Below all functions related to the FizzbuzzHandler

// Function isConformed to check if the request is conformed to the schema expected (five parameters : three integers int1, int2 and limit, and two strings str1 and str2.) 
// and to fill a FizzbuzzParams struct with queries defined
func isConformed(r *http.Request) (p FizzbuzzParams, err error){
	var nbValues int

	err = r.ParseForm()
	if err != nil {
		return p, err
	}

	// Check the URL query
	for key := range r.Form {
		// We increment nbValues to get the number of total parameters of the query
		nbValues += 1

		// We check if keys match with expected keys and we store the associated value in a fizzbuzzParams struct.
		switch key {
		case "int1": {
			// We check if the string value is an integer and if ok we convert the string value to an integer and store it to the fizzbuzzParams struct
			p.Int1, err = strconv.Atoi(r.FormValue("int1"))
			if err != nil {
				return p, err
			}
		}
		case "int2": {
			p.Int2, err = strconv.Atoi(r.FormValue("int2"))
			if err != nil {
				return p, err
			}
		}
		case "limit": {
			p.Limit, err = strconv.Atoi(r.FormValue("limit"))
			if err != nil {
				return p, err
			}		
		}
		case "str1": {
			p.Str1 = r.FormValue("str1")
		}
		case "str2":
			p.Str2 = r.FormValue("str2")
		default: {
			// The key used by the user is not expected.
			err := errors.New("Bad parameter used\nPlease use these 5 keys: int1, int2, limit, str1, str2\nExample: http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz\n\n")
			return p, err
		}
		}
	}	
	
	if nbValues == 0 {
		// Return error if there is no parameter
		err := errors.New("To use Fizzbuzz, please add parameters to URL like this : http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz")
		return p, err
	} else if nbValues != 5 {
		// Return error if there are not 5 parameters or if the same key is used many times
		err := errors.New("Syntax error\nPlease use these 5 keys: int1, int2, limit, str1, str2\nExample: http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz\n")
		return p, err
	} 
	return p, err
}

// doFizzbuzz is the core function which return the list with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
func doFizzbuzz(p FizzbuzzParams) (fizzbuzzList string){
	for i := 1; i <= p.Limit; i++ {
		if i != 1 {
			fizzbuzzList = fizzbuzzList + ","
		} else if i == 1 {
			fizzbuzzList =  fizzbuzzList + "\""
		}
		switch {
		case i%p.Int1 == 0 && i%p.Int2 == 0: {
			fizzbuzzList = fizzbuzzList + p.Str1 + p.Str2
		}
		case i%p.Int1 == 0:
			fizzbuzzList = fizzbuzzList + p.Str1
		case i%p.Int2 == 0:
			fizzbuzzList = fizzbuzzList + p.Str2
		default:
			fizzbuzzList = fizzbuzzList + strconv.Itoa(i)
		}
		if i == p.Limit {
			fizzbuzzList = fizzbuzzList + "\".\n"
		}
	}
	return fizzbuzzList
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	
	// Check the conformity of the request
	p, err := isConformed(r)
	if err != nil{
		fmt.Fprintf(w, "Syntax Error\nPlease, use numbers for int1, int2 and limit values.\nError details below:\n%s", err)
		return
	}

	// Fizzbuzzlist generation
	fizzbuzzList := doFizzbuzz(p)

	// Print parameters defined in the URL 
	fmt.Fprintf(w, "Parameters defined:\n Int1: %d\n Int2: %d\n Limit: %d\n Str1: %s\n Str2: %s\n\n", p.Int1, p.Int2, p.Limit, p.Str1, p.Str2)
	
	// Print the fizzbuzzlist
	fmt.Fprintf(w, "Result:\n%s\n", fizzbuzzList)
}
