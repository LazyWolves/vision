package hostInfoHandler

import (
	"github.com/shirou/gopsutil/host"
)

func GetHostInfo() (*host.InfoStat, error) {
	return host.Info()
}
