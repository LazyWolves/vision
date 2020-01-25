package models

type SystemdHolder struct {
	ServiceName string
	ServiceState string
}

type SystemdResponseHolder struct {
	Services []SystemdHolder
	NumServices int
	Timestamp int64
	TimestampUTC string
}
