package systemdDriver

import (
	"vision/core/systemdHandler"
	"vision/core/models"
)

func ListSystemdServices(filterBy []string) (*[]models.SystemdHolder, error) {

	return systemdHandler.ListSystemdServices(filterBy)
}

func StartSystemdService(target string) (string, error) {

	return systemdHandler.StartSystemdService(target)
}

func StopSystemdService(target string) (string, error) {

	return systemdHandler.StopSystemdService(target)
}
