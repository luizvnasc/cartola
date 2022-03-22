package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Level:     logrus.DebugLevel,
	Formatter: &logrus.TextFormatter{},
}

func Log() *logrus.Logger {
	return log
}
