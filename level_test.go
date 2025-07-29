package golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("Hello Trace")
	logger.Debug("Hello Debug")
	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")

	// logger.Fatal("Hello Fatal")
	// logger.Panic("Hello Panic")

	fmt.Println("HIIIIIIIIIIIIIIIIIII");
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Hello Trace")
	logger.Debug("Hello Debug")
	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")

	// logger.Fatal("Hello Fatal")
	// logger.Panic("Hello Panic")

	fmt.Println("HIIIIIIIIIIIIIIIIIII");
}