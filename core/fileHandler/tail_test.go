package fileHandler

import (
	"testing"
	"strings"
)

func TestTail(t *testing.T) {

	// setup test file
	createTestFile()

	content, err := ReadFromTail("./test-file.txt", "", "", 5, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	lines := strings.Split(content, "\n")
	numLines := len(lines) - 1

	if numLines != 5 {
		t.Errorf("Number of lines = %d, Expected 5", numLines)
	}

	if !(lines[4] == "read form desired files.") {
		t.Errorf("Expected 'read form desired files' at last line, however not found")
	}

	content, err = ReadFromHead("./test-file.txt", "", "", 100, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	lines = strings.Split(content, "\n")
	numLines = len(lines) - 1

	if numLines != 17 {
		t.Errorf("Number of lines = %d, Expected 17", numLines)
	}

	// remove test file
	removeTestFile()
}
