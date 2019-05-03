package fileDriver

import (
	"vision/core/fileHandler"
	"vision/core/models"
)

func fileDriver(request models.QueryHolder) (string, error) {
	isClean, err := request.Sanitise()
	if err != nil {
		return "", err
	}
}
