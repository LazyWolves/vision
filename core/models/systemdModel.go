package models

type SystemdHolder struct {
	ServiceName string
	ServiceState string
}

type ListSystemdResponseHolder struct {
	Services []SystemdHolder
	NumServices int
	Timestamp int64
	TimestampUTC string
}

type OperateSytemdResponseHolder struct {
	Status string
	Timestamp int64
	TimestampUTC string
}
