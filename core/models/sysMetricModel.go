package models

type SystemMetricQueryHolder struct {

	// type of system metric wanted :  can be cpu, memory as of now
	metricType string
}

type CPUMetric struct {
	loadAvg CPULoadAvgMetric	
}

type CPULoadAvgMetric struct {
	Load1 float64
	Load5 float64
	Load15 float64
}

type MemoryMetric struct {
	VirtualMemory VirtualMemoryMetric
}

type VirtualMemoryMetric struct {
	MemTotal float64
	MemFree float64
	UsedPercent float64
}
