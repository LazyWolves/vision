// Package containing method to process resources
package fileHandler

import (
	"os"
	"io"
	"strings"
    "vision/core/util"
    "github.com/sirupsen/logrus"
)

// This function takes a string of lines. It splits the string into a list
// of lines. Then it filetrs the list to get the desired lines. Once it
// has the desired list of lines it joins them into a single string and
// returns it back.
// Params:
//      lines: String of extracted lines
//      posRegex : regex for filtering
//      negRegex : regex for excluding
func getfilteredLines(lines, posRegex, negRegex string, logger *logrus.Logger) (string) {
	lineList := strings.Split(lines, "\n")
	filteredLines := make([]string, 0, 1)

	for _, line := range lineList {
		if util.CheckPattern(line, posRegex, negRegex) {
			filteredLines =  append(filteredLines, line)
		}
	}
	allLines := strings.Join(filteredLines, "\n")
    if len(filteredLines) == 0 {
        return ""
    }
    if allLines[len(allLines) - 1] == 0 {
        allLines = string(allLines[0: len(allLines) - 1])
    }
	return allLines
}

// This function will read a given resource starting from tail uptil a given numer
// of lines as specified in limit
// Params:
//      path : Path to the resource on filesystem
//      posRegex : The regex to filter lines
//      negregex : The regex to exclude lines
//      numLines : the number of lines to limit to
func ReadFromTail(path, posRegex, negRegex  string, numLines int64, logger *logrus.Logger) (string, error) {
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
        return getfilteredLines(string(b), posRegex, negRegex, logger), nil
    } else if err != nil {
        return "", err
    }

    return "**No error but no text read.**", nil
}
