package models

import (
	"regexp"
)

type ProcDescriptionShort struct {
	Pid int32
	Name string
	CmdLine string
}

type ProcDescriptionLong struct {
	Pid int32
	Ppid int32
	Name string
	CmdLine string
	ExePath string
	Cwd string
	Status string
	Uids []int32
	Gids []int32
	Nice []int32
	NumThreads []int32
}

func (p *ProcDescriptionShort) Filter(filterBy, regex string) (bool, error) {

	matchFound := false

	switch filterBy {
	case "name":
		matchFound, _ = regexp.MatchString(regex, p.Name)
	case "cmdline":
		matchFound, _ = regexp.MatchString(regex, p.CmdLine)
	}

	return matchFound, nil
}
