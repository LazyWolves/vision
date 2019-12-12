package util

import (
	"testing"
	"vision/core/models"
)

func TestAclsAllowAll(t *testing.T) {

	// Creating test config
	testConfig := models.ConfigModel {
		AllowAll: true,
		BlockFor: []string{"/blocktest/", "/blocktest2/subfolder/"},
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

	testPath = "/blocktest2/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err != nil {
		t.Errorf("Expected nil, however found error")
	}

	testPath = "/blocktest2/subfolder/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err == nil {
		t.Errorf("Expected Error, however found nil")
	}
}

func TestAclsBlockAll(t *testing.T) {

	// Creating test config
	testConfig := models.ConfigModel {
		AllowAll: false,
		AllowFor: []string{"/allowtest/", "/allowtest2/subfolder"},
	}

	// Create test path
	testPath := "/somefolder/test.log"

	err := CheckAcls(testPath, &testConfig)

	if err == nil {
		t.Errorf("Expected Error, however found nil")
	}

	testPath = "/allowtest/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err != nil {
		t.Errorf("Expected nil, however found error")
	}

	testPath = "/allowtest2/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err == nil {
		t.Errorf("Expected Error, however found nil")
	}

	testPath = "/allowtest2/subfolder/test.log"

	err = CheckAcls(testPath, &testConfig)

	if err != nil {
		t.Errorf("Expected nil, however found error")
	}
}
