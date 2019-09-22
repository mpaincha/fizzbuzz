# leboncoin
Write a simple fizz-buzz REST server
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

Your goal is to implement a web server that will expose a REST API endpoint that: 
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

-------------

Packages to install:

1 - Package Gorilla Mux
import "github.com/gorilla/mux"

>> install command: go get -u github.com/gorilla/mux

readme.md https://github.com/gorilla/mux


2 - Testify Assert
import "testing"
import "github.com/stretchr/testify/assert"

>> install command:  go get github.com/stretchr/testify