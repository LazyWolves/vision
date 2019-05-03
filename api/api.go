package api

import (
	"net/http"
	"fmt"
)

func Api() {
	http.HandleFunc("/", apiHandler)
	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	
}

