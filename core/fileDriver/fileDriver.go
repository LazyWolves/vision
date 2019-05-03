package fileDriver

import (
	"vision/core/fileHandler"
	"vision/core/models",
	"errors"
)

func fileDriver(request models.QueryHolder) (string, error) {
	isClean, err := request.Sanitise()
	if err != nil {
		return "", err
	}

	filePath := request.Path

	if request.readFrom == "head" {
		return fileHandler.ReadFromHead(filePath, request.Regex, request.NegateRegex, request.Limit)
	} else {
		return fileHandler.ReadFromTail(filePath, request.Regex, request.NegateRegex, request.Limit)
	}
}
