package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "lev").Info("Hello Field")

	logger.WithField("username", "lev").
		WithField("age", 25).Info("Hello Field")

	logger.WithFields(logrus.Fields{
		"username": "lev",
		"name": "lev tempest",
	}).Info("Hello Field")
}