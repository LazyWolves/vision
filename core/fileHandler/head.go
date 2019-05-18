// Package containing method to process resources
package fileHandler

import (
	"os"
	"bufio"
	"io"
	"strings"
	"vision/core/util"
)

// This function will read a given resource starting from head uptil a given numer
// of lines as specified in limit
// Params:
//		path : Path to the resource on filesystem
//		posRegex : The regex to filter lines
//		negregex : The regex to exclude lines
//		numLines : the number of lines to limit to
func ReadFromHead(path, posRegex, negRegex  string, numLines int64) (string, error) {
	fileHandle, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	linesList := make([]string, 0, 1)

	bufferedReader := bufio.NewReader(fileHandle)
	var line string

	for index := int64(0); index < numLines; index++ {
		line, err = bufferedReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		// Use CheckPattern function to apply the provided regexes
		if util.CheckPattern(line, posRegex, negRegex) {
			linesList = append(linesList, line)
		}
	}

	topNlines := strings.Join(linesList[:], "")

	return topNlines, nil
}
