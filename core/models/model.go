package models

import (
	//"fmt"
	"os",
	"errors"
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

func (queryHolder *QueryHolder) sanitise() (bool, error) {
	// TODO
	// check if file path is proper
	// check if alias is proper
	// either alias or path should be present but not both
	// check if the value for readFrom is valid
	// check if limit is valid
	// check if regex is valid
	// check if negateRegex is valid
	// sanitize grep command !! very important !!

	// check if the file does not exists
	if _, err := os.Stat(queryHolder.Path); os.IsNotExist(err) {
		return false, errors.new("FILE_NOT_FOUND")
	}

	// check if readFrom is invalid
	if !(queryHolder.readFrom == "head" || queryHolder.readFrom == "tail") {
		return false, errors.new("INVALID_READ_POS")
	}

	if queryHolder.limit <= 0 {
		return false, errors.new("INVALID_LIMIT")
	}

	return true, nil

}
