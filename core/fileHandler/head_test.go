package fileHandler

import (
	"testing"
	"os"
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

	// remove test file
	removeTestFile()
}
