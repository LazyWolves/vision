package fileDriver

import (
	"vision/core/fileHandler"
	"vision/core/models"
	"vision/core/util"
)

func FileDriver(request *models.QueryHolder, aliases map[string]string, configJson *models.ConfigModel) (string, error) {
	isClean, err := request.Sanitise(aliases)
	if err != nil || !isClean {
		return "", err
	}

	filePath := ""

	if request.Path != "" {
		filePath = request.Path
	} else if request.Alias != "" {
		filePath = aliases[request.Alias]
	}

	errAcl := util.checkAcls(filePath, configJson)
	if errAcl == nil {
		return "", errAcl
	}

	if request.ReadFrom == "head" {
		return fileHandler.ReadFromHead(filePath, request.Regex, request.NegateRegex, request.Limit)
	} else {
		return fileHandler.ReadFromTail(filePath, request.Regex, request.NegateRegex, request.Limit)
	}
}
