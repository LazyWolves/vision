package sysMetricDriver

import (
	"vision/core/models"
	"vision/core/sysMetrichandler"
)

func GetSystemMetrics() *models.SystemMetrics {

	return sysMetrichandler.GetSystemMetrics()
}
