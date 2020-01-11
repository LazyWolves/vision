package procHandler

import (
	"github.com/shirou/gopsutil/process"
	"vision/core/models"
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

func describeProc(pid int32) (*models.ProcDescriptionLong, error) {

	proc, err := process.NewProcess(pid)

	if err != nil {
		return nil, err
	}

	procName, _ := proc.Name()
	procCmdLine, _ := proc.Cmdline()
	procPpid, _ := proc.Ppid()
	procExePath, _ := proc.Exe()
	procCwd, _ := proc.Cwd()
	procStatus, _ := proc.Status()
	procUids, _ := proc.Uids()
	procGids, _ := proc.Gids()
	procNice, _ := proc.Nice()
	procNumThreads, _ := proc.NumThreads()

	procDescription := &models.ProcDescriptionLong{
		Pid: proc.Pid,
		Name: procName,
		CmdLine: procCmdLine,
		Ppid: procPpid,
		ExePath: procExePath,
		Cwd: procCwd,
		Status: procStatus,
		Uids: procUids,
		Gids: procGids,
		Nice: procNice,
		NumThreads: procNumThreads,
	}

	return procDescription, nil

}
