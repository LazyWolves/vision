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
