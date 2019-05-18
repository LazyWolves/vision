// This package contains some utility functions for vision
package util

import (
	"regexp"
)

// Fucntion to check for patters. This fuction is used for filtering lines/texts
// Param :
//		line: string containing the line
//		posRegex : Regex containing patterns to filter desired texts
//		negRegex : Regex containing patterms to exlude desired texts
func CheckPattern(line, posRegex, negRegex string) (bool) {
	if posRegex != "" {
		match, _ := regexp.MatchString(posRegex, line)
		if !match {
			return false
		}
	}

	if negRegex != "" {
		match, _ := regexp.MatchString(negRegex, line)
		if match {
			return false
		}
	}

	return true
}
