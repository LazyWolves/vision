package sysMetrichandler

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"vision/core/models"
)

func getCPUMetrics(CPUMetrics *models.CPUMetrics) error {

	cpuLoad, error := load.Avg()

	if error != nil {

		CPULoadAvgMetrics := models.CPULoadAvgMetrics{
			Load1: cpuLoad.Load1,
			Load5: cpuLoad.Load5,
			Load15: cpuLoad.Load15,
		}

		CPUMetrics.LoadAvg = CPULoadAvgMetrics

		return nil
	}

	return error
}
