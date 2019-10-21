package fileHandler

import (
	"testing"
	"os"
	"strings"
)

func createTestFile() {

	// Content for test file
	fileContent := `This is a test file for performing unit testing
on the functions provided in this package. Before
each test function runs, this content will be dumped
to a file named test-file.txt. Once the test case has
been executed, this file will be removed.
Vision is a light weight tool written purely in
golang for viewing remote resources over HTTP.
Vision allows you to view
config files, log files and other such
files over HTTP via your browser or on your terminal.
It allows you to set ACLs via
which you can block view on certain resources and
alow view on certain resources. It allows you toconfigure aliases
so that you do not have to type the entire path of the resource
on server, view a file from top, or bottom, apply regex
for filtering contents and specify number of lines to be
read form desired files.
`

	f, _ := os.Create("test-file.txt")
	f.WriteString(fileContent)
	f.Close()
}

func removeTestFile() {
	
	//Remove test-file.txt
	os.Remove("test-file.txt")
}

func TestHead(t *testing.T) {

	// setup test file
	createTestFile()

	content, err := ReadFromHead("./test-file.txt", "", "", 5, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	lines := strings.Split(content, "\n")
	numLines := len(lines) - 1

	if numLines != 5 {
		t.Errorf("Number of lines = %d, Expected 5", numLines)
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

func TestHeadWithPosFiler(t *testing.T) {

	// setting up test file
	createTestFile()

	content, err := ReadFromHead("./test-file.txt", "HTTP", "", 100, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	lines := strings.Split(content, "\n")
	numLines := len(lines) - 1

	if numLines != 2 {
		t.Errorf("Number of lines = %d, Expected 1", numLines)
	}

	if !strings.Contains(lines[0], "HTTP") {
		t.Errorf("Expected HTTP, however not found")
	}

	// removing test file
	removeTestFile()
}

func TestHeadWithNegFiler(t *testing.T) {

	// setting up test file
	createTestFile()

	content, err := ReadFromHead("./test-file.txt", "", "HTTP", 100, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	lines := strings.Split(content, "\n")
	numLines := len(lines) - 1

	if numLines != 15 {
		t.Errorf("Number of lines = %d, Expected 15", numLines)
	}

	linesConcatinated := strings.Join(lines, "")
	if strings.Contains(linesConcatinated, "HTTP") {
		t.Errorf("HTTP not expected, however got it")
	}

	// removing test file
	removeTestFile()
}

func TestHeadWithRegexFiler(t *testing.T) {

	// setting up test file
	createTestFile()

	content, err := ReadFromHead("./test-file.txt", "HTTP|ACLs", "", 100, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	lines := strings.Split(content, "\n")
	numLines := len(lines) - 1

	if numLines != 3 {
		t.Errorf("Number of lines = %d, Expected 3", numLines)
	}

	linesConcatinated := strings.Join(lines, "")
	t.Logf(linesConcatinated)
	if !(strings.Contains(linesConcatinated, "HTTP") && strings.Contains(linesConcatinated, "ACLs")) {
		t.Errorf("HTTP and ACLs expected, however not found")
	}

	// removing test file
	removeTestFile()
}
