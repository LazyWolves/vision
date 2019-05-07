package fileDriver

import (
	"vision/core/fileHandler"
	"vision/core/models"
)

func FileDriver(request *models.QueryHolder, aliases map[string]string) (string, error) {
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

	if request.ReadFrom == "head" {
		return fileHandler.ReadFromHead(filePath, request.Regex, request.NegateRegex, request.Limit)
	} else {
		return fileHandler.ReadFromTail(filePath, request.Regex, request.NegateRegex, request.Limit)
	}
}
