package log4

import "github.com/sirupsen/logrus"

// Version information
const (
	VERSION = "v0.1.0"
	MAJOR   = 0
	MINOR   = 1
	BUILD   = 0
)

type Logger4 interface {
	// Fatal(arg0 interface{}, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	// Debug(arg0 interface{}, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	// Info(arg0 interface{}, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	// Warning(arg0 interface{}, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	// Error(arg0 interface{}, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry
}
