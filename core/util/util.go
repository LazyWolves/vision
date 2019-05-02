package util

import (
	"regexp"
)

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
