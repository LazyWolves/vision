package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Function to setup logging using using provided file handler
func SetupLogger(fileHandle *os.File) *logrus.Logger {

	// Get a new instance of logrus. An instances is being created so
	// that it can be passed among functions
	logger := logrus.New()
	return logger
}

