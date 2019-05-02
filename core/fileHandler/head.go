package fileHandler

import (
	"os"
	//"errors"
	"bufio"
	"io"
	//"fmt"
	"strings"
	"vision/core/util"
)

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
		if util.CheckPattern(line, posRegex, negRegex) {
			linesList =  append(linesList, line)
		}
	}

	topNlines := strings.Join(linesList[:], "")

	return topNlines, nil
}
