package sysMetricDriver

import (
	"vision/core/models"
	"vision/core/sysMetrichandler"
)

func getSystemMetrics() *models.SystemMetrics {

	return sysMetrichandler.GetSystemMetrics()
}
