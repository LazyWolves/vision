package models

import (
	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	HostInfo host.InfoStat
	Timestamp int64
	TimestampUTC string
}
