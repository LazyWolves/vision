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

func isValidPath(path string) (error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("FILE_NOT_FOUND")
	}

	if err != nil {
		return err
	}

	if stat.IsDir() {
		return errors.New("PATH_IS_A_DIRECTORY")
	}

	return nil
}

func (queryHolder *QueryHolder) Sanitise(aliases [string]string) (bool, error) {
	// TODO
	// check if regex is valid
	// check if negateRegex is valid
	// sanitize grep command !! very important !!
	path := ""

	if queryHolder.Path != "" {
		path = queryHolder.Path
	} else if queryHolder.Alias != "" {
		path = aliases[queryHolder.Alias]
	} else {
		return false, errors.New("BOTH_PATH_AND_ALIAS_IS_EMPTY")
	}

	err := isValidPath(path)
	if err != nil {
		return false, err
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
