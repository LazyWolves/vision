// Package containing struct definitions for holding configs and params used
// in vision
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

// Method to check if the given path is valid. The resource pointed by the
// path should exists and the resource should not be a directory
// Params :
//		path : atring containing resource
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

// Function to sanitise the URL query params
// Params:
//		aliases : alias map
func (queryHolder *QueryHolder) Sanitise(aliases map[string]string) (bool, error) {
	path := ""
	exists := false

	// Both path and alias cannot be defined together
	// Also path will be given higher priority over alias
	if queryHolder.Path != "" {
		path = queryHolder.Path
	} else if queryHolder.Alias != "" {
		path, exists = aliases[queryHolder.Alias]
		if !exists {
			return false, errors.New("ALIAS_DOES_NOT_EXISTS")
		}
	} else {
		return false, errors.New("BOTH_PATH_AND_ALIAS_IS_EMPTY")
	}

	// checks if path is valid.
	err := isValidPath(path)
	if err != nil {
		return false, err
	}

	// check if readFrom is invalid
	if !(queryHolder.ReadFrom == "head" || queryHolder.ReadFrom == "tail") {
		return false, errors.New("INVALID_READ_POS")
	}

	// limit cannot be negative or 0
	if queryHolder.Limit <= 0 {
		return false, errors.New("INVALID_LIMIT")
	}

	return true, nil
}
