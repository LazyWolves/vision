package util

import (
	"github.com/sirupsen/logrus"
)

// Function to setup logging using using provided file handler
func SetupLogger() *logrus.Logger {

	// Get a new instance of logrus. An instances is being created so
	// that it can be passed among functions
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	return logger
}

