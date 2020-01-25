package systemdHandler

import (
	"vision/core/models"
	"github.com/coreos/go-systemd/dbus"
)

func ListSystemdServices(filterBy []string) (*[]models.SystemdHolder, error) {

	listSystemdServices := []models.SystemdHolder{}

	sbusConn, err := dbus.New()
	
	if err != nil {
		return &listSystemdServices, err
	}

	units, err := sbusConn.ListUnitsByPatterns([]string{}, filterBy)

	if err != nil {
		return &listSystemdServices, err
	}

	for _, unit := range units {
		listSystemdServices = append(listSystemdServices, models.SystemdHolder{
			ServiceName: unit.Name,
			ServiceState: unit.ActiveState,
		})
	}

	return &listSystemdServices, nil
}
