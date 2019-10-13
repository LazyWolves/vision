// This package generates response for a iven request
package fileDriver

import (
	"vision/core/fileHandler"
	"vision/core/models"
	"vision/core/util"
	"github.com/sirupsen/logrus"
)

// This function processes the requests and generates the response. It uses sanitise function
// to sanitise the requests and then uses fileHandler package to generate response.
// Params:
// 		request : Struct of type QuireyHolder containing all the URL params
//		aliases : Alias map
//		configJson : Struct of type ConfigModel containing all the config params
func FileDriver(request *models.QueryHolder, aliases map[string]string, configJson *models.ConfigModel, logger *logrus.Logger) (string, error) {

	// Sanitise the request
	isClean, err := request.Sanitise(aliases)
	if err != nil || !isClean {
		return "", err
	}

	filePath := ""

	// Set resource path, path param is given more preference than alias param.
	// So if both path and alias are there in URL params, path is chosen.
	if request.Path != "" {
		filePath = request.Path
	} else if request.Alias != "" {
		filePath = aliases[request.Alias]
	}

	// Evaluate Acls and check it the current resource is allowed to be viewed.
	// If allowed then proceed, if not then send suitable message back to user
	errAcl := util.CheckAcls(filePath, configJson)
	if errAcl != nil {
		logger.Error("Access right violation for path : ", filePath)
		return "", errAcl
	}

	// Read from head or from tail as the request may be.
	if request.ReadFrom == "head" {
		return fileHandler.ReadFromHead(filePath, request.Regex, request.NegateRegex, request.Limit, logger)
	} else {
		return fileHandler.ReadFromTail(filePath, request.Regex, request.NegateRegex, request.Limit, logger)
	}
}
