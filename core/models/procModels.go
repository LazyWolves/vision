package models

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
}
