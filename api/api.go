// api package contains the main api endpoints which allows users to view resources
package api

import (
	"net/http"
	"fmt"
	"strconv"
	"vision/core/fileDriver"
	"vision/core/models"
	"strings"
	"io/ioutil"
	"encoding/json"
)

// configJson struct is the model for storing and holding the config data
var configJson models.ConfigModel

// aliases is a hashmap to create a one to one mapping
// between the alias name and resource path
var aliases map[string]string

// Store path to config file
var configJsonPath = "/etc/vision/config.json"

// Main function which loads config, creates alias hash and attaches handlers to routes
func Api() {
	loadConfigJson()
	createAliasMap()
	http.HandleFunc("/", apiHandler)
	http.HandleFunc("/aliases", aliasHandler)
	http.ListenAndServe(":" + strconv.FormatInt(configJson.Port, 10), nil)
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

func aliasHandler(w http.ResponseWriter, r *http.Request) {
	response := allAliases()
	fmt.Fprintf(w, response)
}

func allAliases() (string) {
	aliasesSlice := make([]string, 0, 10)
	for key, value := range aliases {
		aliasesSlice = append(aliasesSlice, key + " : " + value)
	}

	if len(aliasesSlice) != 0 {
		aliasesString := strings.Join(aliasesSlice, "\n")
		return aliasesString
	}

	return ""
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

	response, err := fileDriver.FileDriver(request, aliases, &configJson)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, response)
}
