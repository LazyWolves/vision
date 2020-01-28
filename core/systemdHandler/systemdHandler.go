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

func isValidService(service string) (bool, error) {

	sbusConn, err := dbus.New()
	units, err := sbusConn.ListUnitsByPatterns([]string{}, []string{service})

	if err != nil {
		return false, err
	}

	for _, unit := range units {
		if unit.Name == service {
			return true, nil
		}
	}

	return false, nil
}

func StartSystemdService(target string) (string, error) {

	target = target + ".service"

	sbusConn, err := dbus.New()
	if err != nil {
		return "FAILED", err
	}

	servicePresent, err := isValidService(target)

	if err != nil {
		return "FAILED", err
	}

	if !servicePresent {
		return "SERVICE DOES NOT EXIST", err
	}

	resChan := make(chan string)

	_, err = sbusConn.StartUnit(target, "replace", resChan)

	if err != nil {
		return "FAILED", err
	}

	job := <- resChan

	if job != "done" {
		return "FAILED", nil
	}

	return "OK", nil
}

func StopSystemdService(target string) (string, error) {

	target = target + ".service"

	sbusConn, err := dbus.New()
	if err != nil {
		return "FAILED", err
	}

	servicePresent, err := isValidService(target)

	if err != nil {
		return "FAILED", err
	}

	if !servicePresent {
		return "SERVICE DOES NOT EXIST", err
	}

	resChan := make(chan string)

	_, err = sbusConn.StopUnit(target, "replace", resChan)

	if err != nil {
		return "FAILED", err
	}

	_ = <- resChan

	return "OK", nil
}
