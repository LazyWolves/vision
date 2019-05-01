package fileHandler

import (
	"os"
	"errors"
	"bufio"
	"io"
	"fmt"
)

func ReadFromHead(path string, numLines int64) (string, error) {
	fileHandle, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fileHandle.close()

	bufferedReader := bufio.NewReader(fileHandle)
	var line string

	for index := 0; index < numLines; index++ {
		line, err = bufferedReader.readString("\n")
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		fmt.Println(line)
	}

	return line, nil
}
