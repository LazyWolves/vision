package procHandler

import (
	"github.com/shirou/gopsutil/process"
	"vision/core/models"
	"fmt"
)

func listAllProcs() (*[]models.ProcDescriptionShort, error) {

	procList := make([]models.ProcDescriptionShort, 1) 
	processes, _ := process.Processes()

	for _, proc := range processes {
		
		procName, _ := proc.Name()
		procCmdLine, _ := proc.Cmdline()

		procHolder := models.ProcDescriptionShort{
			Pid: proc.Pid,
			Name: procName,
			CmdLine: procCmdLine,
		}

		procList = append(procList, procHolder)
	}

	return &procList, nil
}
