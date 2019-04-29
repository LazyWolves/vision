package models

import (
	//"fmt"
)

type QueryHolder struct {
	// path to the file
	Path string

	// alias name for a file. must be configured
	alias string

	// can accept only two values : head or tail
	readFrom string

	// number of lines to be streamed
	limit int32

	// search for lines containing entities matching given regex
	regex string

	// search for lines containing entities which does not match given regex
	negateRegex string

	// grep command options
	grep string
}

func (queryHolder *QueryHolder) sanitise() {
	// TODO
	// check if file path is proper
	// check if alias is proper
	// either alias or path should be present but not both
	// check if the value for readFrom is valid
	// check if limit is valid
	// check if regex is valid
	// check if negateRegex is valid
	// sanitize grep command !! very important !!
}
