package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	// Importable global logrus logger instance.
	Logger *logrus.Logger = logrus.New()

	// DebugLevel level. Usually only enabled when debugging.
	// Very verbose logging. (logrus.Level = 5)
	Debug logrus.Level = logrus.DebugLevel

	// InfoLevel level. General operational entries about what's
	// going on inside the application. (logrus.Level = 4)
	Info logrus.Level = logrus.InfoLevel
)
