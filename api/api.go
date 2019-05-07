package api

import (
	"net/http"
	"fmt"
	"strconv"
	//"reflect"
	"vision/core/fileDriver"
	"vision/core/models"
	//"os"
	"io/ioutil"
	"encoding/json"
)

var configJson models.ConfigModel
var aliases map[string]string
var configJsonPath = "/etc/vision/config.json"

func Api() {
	loadConfigJson()
	createAliasMap()
	http.HandleFunc("/", apiHandler)
	http.ListenAndServe(":8080", nil)
}

func loadConfigJson() {
	file, _ := ioutil.ReadFile(configJsonPath)
	_ = json.Unmarshal([]byte(file), &configJson)
}

func createAliasMap() {
	aliasesTemp := make(map[string]string)
	for _, alias := range configJson.Aliases {
		aliasesTemp[alias.AliasName] = alias.AliasTo
	}
	aliases = aliasesTemp
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	pathSlice, isPath := r.URL.Query()["path"]
	readFromSlice, isReadFrom := r.URL.Query()["readFrom"]
	limitSlice, isLimit := r.URL.Query()["limit"]
	posRegexSlice, isPosRegex := r.URL.Query()["filterBy"]
	negRegexSlice, isNegRegex := r.URL.Query()["ignore"]
	aliasSlice, isAlias := r.URL.Query()["alias"]

	path, readFrom, limit, posRegex, negRegex, alias := "", "tail", int64(10), "", "", ""

	if isPath {
		path = pathSlice[0]
	}

	if isReadFrom {
		readFrom = readFromSlice[0]
	}

	if isLimit {
		limitTemp, err := strconv.ParseInt(limitSlice[0], 10, 64)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		limit = limitTemp
	}

	if isPosRegex {
		posRegex = posRegexSlice[0]
	}

	if isNegRegex {
		negRegex = negRegexSlice[0]
	}

	if isAlias {
		alias = aliasSlice[0]
	}

	request := &models.QueryHolder{
		Path: path,
		Alias: alias,
		ReadFrom: readFrom,
		Limit: limit,
		Regex: posRegex,
		NegateRegex: negRegex,
		Grep: "",
	}

	//fmt.Println(path)
	//fmt.Println(readFrom)
	//fmt.Println(limit)
	//fmt.Println(posRegex)
	//fmt.Println(negRegex)

	response, err := fileDriver.FileDriver(request, aliases)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, response)
}
