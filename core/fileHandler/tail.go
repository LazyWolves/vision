package fileHandler

import (
	"os"
	//"errors"
	//"bufio"
	"io"
	"strings"
	"vision/core/util"
)

func getfilteredLines(lines, posRegex, negRegex string) (string) {
	lineList := strings.Split(lines, "\n")
	filteredLines := make([]string, 0, 1)

	for _, line := range lineList {
		if util.CheckPattern(line, posRegex, negRegex) {
			filteredLines =  append(filteredLines, line)
		}
	}

	allLines := strings.Join(filteredLines, "\n")
	return allLines
}

func ReadFromTail(path, posRegex, negRegex  string, numLines int64) (string, error) {
    fileHandle, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer fileHandle.Close()

    numNewLines := int64(0)
    var offset int64 = -1
    var finalReadStartPos int64
    for numNewLines <= numLines-1 {
        startPos, err := fileHandle.Seek(offset, 2)
        if err != nil {
            return "", err
        }

        if startPos == 0 {
            finalReadStartPos = -1
            break
        }

        b := make([]byte, 1)
        _, err = fileHandle.ReadAt(b, startPos)
        if err != nil {
            return "", err
        }

        if offset == int64(-1) && string(b) == "\n" {
            offset--
            continue
        }

        if string(b) == "\n" {
            numNewLines++
            finalReadStartPos = startPos
        }

        offset--
    }

    endPos, err := fileHandle.Seek(int64(-1), 2)
    if err != nil {
        return "", err
    }
    b := make([]byte, (endPos+1)-finalReadStartPos)
    _, err = fileHandle.ReadAt(b, finalReadStartPos+1)
    if err == io.EOF {
        return getfilteredLines(string(b), posRegex, negRegex), nil
    } else if err != nil {
        return "", err
    }

    return "**No error but no text read.**", nil
}
