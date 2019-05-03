package fileDriver

import (
	"vision/core/fileHandler"
	"vision/core/models"
	"fmt"
	//"errors"
)

func FileDriver(request *models.QueryHolder) (string, error) {
	isClean, err := request.Sanitise()
	fmt.Println(isClean)
	if err != nil {
		return "", err
	}

	filePath := request.Path

	if request.ReadFrom == "head" {
		return fileHandler.ReadFromHead(filePath, request.Regex, request.NegateRegex, request.Limit)
	} else {
		return fileHandler.ReadFromTail(filePath, request.Regex, request.NegateRegex, request.Limit)
	}
}
