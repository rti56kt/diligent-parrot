package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger = logrus.New()
	Debug  logrus.Level   = logrus.DebugLevel
	Info   logrus.Level   = logrus.InfoLevel
)
