package fileHandler

import (
	"testing"
	"os"
)

func createTestFile() {

	// Content for test file
	fileContent := `
		This is a test file for performing unit testing
		on the functions provided in this package. Before
		each test function runs, this content will be dumped
		to a file named test-file.txt. Once the test case has
		been executed, this file will be removed.
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

}
