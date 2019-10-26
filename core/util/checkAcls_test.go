package util

import (
	"testing"
	"vision/core/models"
)

func TestAclsAllowAll(t *testing.T) {

	// Creating test config
	testConfig := models.ConfigModel {
		AllowAll: true,
		BlockFor: []string{"/blocktest, /blocktest2/subfolder"},
	}

	// Create test path
	testPath := "/somefolder/test.log"

	err := CheckAcls(testPath, &testConfig)

	if err != nil {
		t.Errorf("Expected nil, however got error")
	}
}
