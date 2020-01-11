package procDriver

import (
	"vision/core/models"
	"vision/core/procHandler"
)

func GetListOfProcesses(filterBy, regex string) (*[]models.ProcDescriptionShort, error) {

	procList, err := procHandler.ListAllProcs()

	if err != nil {
		return nil, err
	}

	procListFilter := make([]models.ProcDescriptionShort, 1)

	for _, proc := range *procList {
		if filterBy == "" && regex == "" {
			procListFilter = append(procListFilter, proc)
			continue
		}

		wanted, _ := proc.Filter(filterBy, regex)

		if wanted {
			procListFilter = append(procListFilter, proc)
		}
	}

	return &procListFilter, nil
}

func GetProcessDetails(pid int32) (*models.ProcDescriptionLong, error) {

	process, err := procHandler.DescribeProc(pid)

	if err != nil {
		return nil, err
	}

	return process, nil
}
