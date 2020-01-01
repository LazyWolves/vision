package sysMetrichandler

import (
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"vision/core/models"
)

func getCPUMetrics(CPUMetrics *models.CPUMetrics) {

	cpuLoad, error := load.Avg()

	if error == nil {

		CPULoadAvgMetrics := models.CPULoadAvgMetrics{
			Load1: cpuLoad.Load1,
			Load5: cpuLoad.Load5,
			Load15: cpuLoad.Load15,
		}

		CPUMetrics.LoadAvg = CPULoadAvgMetrics
	}
}

func getMemoryMetrics(MemoryMetrics *models.MemoryMetrics) {

	virtualMemory, error := mem.VirtualMemory()

	if error == nil {

		VirtualMemoryMetrics := models.VirtualMemoryMetrics{
			MemTotal: virtualMemory.Total,
			MemFree: virtualMemory.Free,
			UsedPercent: virtualMemory.UsedPercent,
		}

		MemoryMetrics.VirtualMemory = VirtualMemoryMetrics
	}
}

func GetSystemMetrics() *models.SystemMetrics{

	cpuMetrics := models.CPUMetrics{}
	memoryMetrics := models.MemoryMetrics{}

	getCPUMetrics(&cpuMetrics)
	getMemoryMetrics(&memoryMetrics)

	SystemMetrics := models.SystemMetrics{
		CPU: cpuMetrics,
		Memory: memoryMetrics,
	}

	return &SystemMetrics
}
