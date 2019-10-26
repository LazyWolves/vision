package util

import (
	"testing"
	"vision/core/models"
)

func TestAclsAllowAll(t *testing.T) {

	// Creating test config
	testConfig := models.ConfigModel {
		AllowAll: true,
		BlockFor: []string{"/blocktest", "/block2/subfolder"},
	}

	// Create test path
	testPath := "/somefolder/test.log"

	err := CheckAcls(testPath, &testConfig)

	if err != nil {
		t.Errorf("Expected nil, however got error")
	}

	testPath = "/blocktest/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err == nil {
		t.Errorf("Expected error, however found nil")
	}

	testPath = "/block/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err != nil {
		t.Errorf("Expected nil, however found error")
	}

	testPath = "/block2/subfolder/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err == nil {
		t.Errorf("Expected Error, however found nil")
	}
}
