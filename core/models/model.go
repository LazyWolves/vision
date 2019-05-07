package models

import (
	"os"
	"errors"
)

type QueryHolder struct {
	// path to the file
	Path string

	// alias name for a file. must be configured
	Alias string

	// can accept only two values : head or tail
	ReadFrom string

	// number of lines to be streamed
	Limit int64

	// search for lines containing entities matching given regex
	Regex string

	// search for lines containing entities which does not match given regex
	NegateRegex string

	// grep command options
	Grep string
}

func (queryHolder *QueryHolder) Sanitise() (bool, error) {
	// TODO
	// check if alias is proper
	// either alias or path should be present but not both
	// check if the value for readFrom is valid
	// check if regex is valid
	// check if negateRegex is valid
	// sanitize grep command !! very important !!

	// check if the file does not exists or if the path is a directory

	stat, err := os.Stat(queryHolder.Path)
	if os.IsNotExist(err) {
		return false, errors.New("FILE_NOT_FOUND")
	}

	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return false, errors.New("PATH_IS_A_DIRECTORY")
	}

	// check if readFrom is invalid
	if !(queryHolder.ReadFrom == "head" || queryHolder.ReadFrom == "tail") {
		return false, errors.New("INVALID_READ_POS")
	}

	if queryHolder.Limit <= 0 {
		return false, errors.New("INVALID_LIMIT")
	}

	return true, nil
}
