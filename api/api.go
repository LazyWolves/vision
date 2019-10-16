// api package contains the main api endpoints which allows users to view resources
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"vision/core/fileDriver"
	"vision/core/models"
	"vision/core/util"
	"github.com/sirupsen/logrus"
)

// configJson struct is the model for storing and holding the config data
var configJson models.ConfigModel

// Logger object for writing logs to file
var logger *logrus.Logger

// aliases is a hashmap to create a one to one mapping
// between the alias name and resource path
var aliases map[string]string

// Store path to config file
var configJsonPath = "/etc/vision/config.json"

// Store path to log file
var logFilePath = "/var/log/vision/vision.log"

// Main function which loads config, creates alias hash and attaches handlers to routes
func Api() {
	// Read the config file and load it into memory as json
	// The connfig json will be stored in memory throughout the life
	// time of the object for fast retrieval of config
	loadConfigJson()

	// Create the alias map for fast retrieval.
	// The map will be stored in memory all the time
	// to allow repeated file access
	createAliasMap()

	// Open log file for logging purpose and initialise logger object
	fileHandler, err := os.OpenFile(logFilePath,  os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	// CLose the log file before the function returns
	defer fileHandler.Close()

	// Get an instance of logger
	logger = util.SetupLogger()

	// Notify user if log file could not be opened
	if err != nil {
		logger.Warning("Could not open log file : ", logFilePath)
	} else {

		// If the log file could be loaded then use it in logger object
		logger.SetOutput(fileHandler)
	}

	// Create route for / path and add handler function for it
	http.HandleFunc("/", apiHandler)

	// Create route for /aliases path and add handler function to it
	http.HandleFunc("/aliases", aliasHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(configJson.Port, 10), nil))
}

// This function will read config file and load the json into memory
func loadConfigJson() {
	file, err := ioutil.ReadFile(configJsonPath)
	if err != nil {
		fmt.Printf("Config file not found.\n")
		os.Exit(1)
	}
	_ = json.Unmarshal([]byte(file), &configJson)
}

// This function will create a map (using the data in config json) to
// store aliases and their corresponding paths and store it in memory.
func createAliasMap() {
	aliasesTemp := make(map[string]string)
	for _, alias := range configJson.Aliases {
		aliasesTemp[alias.AliasName] = alias.AliasTo
	}
	aliases = aliasesTemp
}

// This is the handler for serving aliases. It returns the
// alias map as a list
func aliasHandler(w http.ResponseWriter, r *http.Request) {
	response := allAliases()
	fmt.Fprintf(w, response)
}

// This function creates a string representing the alias map with
// proper formatting.
func allAliases() string {
	aliasesSlice := make([]string, 0, 10)
	for key, value := range aliases {
		aliasesSlice = append(aliasesSlice, key+" : "+value)
	}

	if len(aliasesSlice) != 0 {
		aliasesString := strings.Join(aliasesSlice, "\n")
		return aliasesString
	}

	return ""
}

// This is the handler for root. It takes in a number of URL query params
// and it returns resource accordingly.
// It takes the followung URL params :
// 		path : The path (absolute) to the resource to be viewed. For example it can be path to
//			   a log file
//		readFrom : It specifies the end from which the resource is to be read -
//				   head or tail. Accordingly it can take only two values: head|tail
//		limit : It denotes the number of lines to be read from that resource.
//				For example the number of lines of a log file to be read.
//				It must be a integer greater than 0
//		filterBy : The value should be a regex. The regex will be used to filter lines
//				   from the resource specified and only those lines will be returned.
//		ignore :   The value should be a regex. The regex will be used as a negative
//				   filter to remove lines which will contain texts matching the regex.
//		alias :	   This represents alias to a path. Must be configured in config json.
func apiHandler(w http.ResponseWriter, r *http.Request) {

	// The URL parameters are extracted and stored in respective variables
	pathSlice, isPath := r.URL.Query()["path"]
	readFromSlice, isReadFrom := r.URL.Query()["readFrom"]
	limitSlice, isLimit := r.URL.Query()["limit"]
	posRegexSlice, isPosRegex := r.URL.Query()["filterBy"]
	negRegexSlice, isNegRegex := r.URL.Query()["ignore"]
	aliasSlice, isAlias := r.URL.Query()["alias"]

	// variable to store request log
	requestLog := make([]string, 0, 1)

	// Get remote client
	remote_client := r.RemoteAddr

	path, readFrom, limit, posRegex, negRegex, alias := "", "tail", int64(10), "", "", ""

	if isPath {
		path = pathSlice[0]
		requestLog = append(requestLog, "path : " + path)
	}

	if isReadFrom {
		readFrom = readFromSlice[0]
		requestLog = append(requestLog, "readFrom : " + readFrom)
	}

	// Convert the limit to int64 type and return error if any during conversion
	if isLimit {
		limitTemp, err := strconv.ParseInt(limitSlice[0], 10, 64)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		limit = limitTemp
		requestLog = append(requestLog, "limit : " + limitSlice[0])
	}

	if isPosRegex {
		posRegex = posRegexSlice[0]
		requestLog = append(requestLog, "filterBy : " + posRegex)
	}

	if isNegRegex {
		negRegex = negRegexSlice[0]
		requestLog = append(requestLog, "ignore : " + negRegex)
	}

	if isAlias {
		alias = aliasSlice[0]
		requestLog = append(requestLog, "alias : " + alias)
	}

	// Store all the URL params in QueryHolder struct.
	// This is for easy handling of the request
	request := &models.QueryHolder{
		Path:        path,
		Alias:       alias,
		ReadFrom:    readFrom,
		Limit:       limit,
		Regex:       posRegex,
		NegateRegex: negRegex,
		Grep:        "",
	}

	// Log request
	logger.WithFields(logrus.Fields{
		"remote_client": remote_client,
	}).Info(strings.Join(requestLog, " "))

	// Get the response for the current request and write it to the response
	// of the current request and send it ot user. If FIleDriver returns
	// any error then send it to user.
	response, err := fileDriver.FileDriver(request, aliases, &configJson, logger)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"remote_client": remote_client,
		}).Error(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}

	// Send response back to user
	fmt.Fprintf(w, response)
}
