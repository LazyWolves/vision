package hostInfoDriver

import (
	"vision/core/hostInfoHandler"
	"github.com/shirou/gopsutil/host"
)

func HostInfo() (*host.InfoStat, error) {
	return hostInfoHandler.GetHostInfo()
}
