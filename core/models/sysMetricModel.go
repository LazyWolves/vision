package models

type SystemMetrics struct {

	// type of system metric wanted : can be cpu, memory as of now
	CPU CPUMetrics
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
	MemTotal uint64
	MemFree uint64
	UsedPercent float64
}
