package api

import (
	"net/http"
	"fmt"
	"strconv"
	//"reflect"
	//"vision/core/fileDriver"
	//"vision/core/models"
)

func Api() {
	http.HandleFunc("/", apiHandler)
	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	pathSlice, isPath := r.URL.Query()["path"]
	readFromSlice, isReadFrom := r.URL.Query()["readFrom"]
	limitSlice, isLimit := r.URL.Query()["limit"]
	posRegexSlice, isPosRegex := r.URL.Query()["posRegexSlice"]
	negRegexSlice, isNegRegex := r.URL.Query()["negRegexSlice"]

	path, readFrom, limit, posRegex, negRegex := "", "head", int64(0), "", ""

	if isPath {
		path = pathSlice[0]
	}

	if isReadFrom {
		readFrom = readFromSlice[0]
	}

	if isLimit {
		limitTemp, err := strconv.ParseInt(limitSlice[0], 10, 64)
		limit = limitTemp
		fmt.Print(err)
	}

	if isPosRegex {
		posRegex = posRegexSlice[0]
	}

	if isNegRegex {
		negRegex = negRegexSlice[0]
	}

	fmt.Println(path)
	fmt.Println(readFrom)
	fmt.Println(limit)
	fmt.Println(posRegex)
	fmt.Println(negRegex)
}

