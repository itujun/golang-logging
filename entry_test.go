package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()

	// otomatis oleh logger
	logger.WithField("username", "lev").Info("Hello Entry") 

	// manual
	entry := logrus.NewEntry(logger)
	entry.WithField("username", "lev")
	entry.Info("Hello Entry")
}