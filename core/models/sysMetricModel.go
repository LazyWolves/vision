package models

type SystemMetrics struct {

	// type of system metric wanted : can be cpu, memory as of now
	Load CPULoadAvgMetrics
	Memory MemoryMetrics
}

type CPUMetrics struct {
	LoadAvg CPULoadAvgMetrics
}

type CPULoadAvgMetrics struct {
	Load1 float64
	Load5 float64
	Load15 float64
}

type MemoryMetrics struct {
	VirtualMemory VirtualMemoryMetrics
}

type VirtualMemoryMetrics struct {
	MemTotal float64
	MemFree float64
	UsedPercent float64
}
