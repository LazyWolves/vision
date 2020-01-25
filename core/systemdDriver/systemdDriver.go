package systemdDriver

import (
	"vision/core/systemdHandler"
	"vision/core/models"
)

func ListSystemdServices(filterBy []string) (*[]models.SystemdHolder, error) {

	return systemdHandler.ListSystemdServices(filterBy)
}
