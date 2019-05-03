package api

import (
	"net/http"
	"fmt"
	"vision/core/fileDriver"
	"vision/core/models"
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

	path, readFrom, limit, posRegex, negRegex := "", "head", 0, "", ""

	if !err {
		return
	}

	path := paths[0]
	fmt.Println(path)
}

