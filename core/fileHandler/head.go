package fileHandler

import (
	"os"
	//"errors"
	"bufio"
	"io"
	//"fmt"
	"strings"
)

func ReadFromHead(path string, numLines int64) (string, error) {
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
		linesList =  append(linesList, line)
	}

	topNlines := strings.Join(linesList[:], "")

	return topNlines, nil
}
