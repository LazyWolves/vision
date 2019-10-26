package util

import (
	"testing"
)

func TestUtilPosRegex(t *testing.T) {

	// Create test string
	testString := "This is a test string for positive regex match. Testing-1,2,3"
	passed := CheckPattern(testString, "regex", "")

	if !passed {
		t.Errorf("Expected true, however got false")
	}

	passed = CheckPattern(testString, "Testing-[0-9]+.*", "")

	if !passed {
		t.Errorf("Expected true, however got false")
	}

	passed = CheckPattern(testString, "^Testing-[0-9]+.*$", "")

	if passed {
		t.Errorf("Expected false, however got true")
	}

	passed = CheckPattern(testString, "absent", "")
}

func TestUtilNegRegex(t *testing.T) {

	// Creating test string
	testString := "This is a test string for positive regex match. Testing-123"
	passed := CheckPattern(testString, "", "regex")

	if passed {
		t.Errorf("Expected false, however got true")
	}

	passed = CheckPattern(testString, "", "Testing-[1-3]+")

	if passed {
		t.Errorf("Expected false, however got true")
	}

	passed = CheckPattern(testString, "", "absent")

	if !passed {
		t.Errorf("Expected true, however got false")
	}
}
