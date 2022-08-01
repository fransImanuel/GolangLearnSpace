package belajargolanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username","frans").Info("Hello Logger")
	
	logger.WithField("username","frans").
		WithField("name","frans imanuel").
		Info("Hello Logger")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username":"Frans",
		"name":"Frans Imanuel",
	}).Infof("Hello Worlds")
}