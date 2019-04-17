package log4

import (
	"os"

	"github.com/sirupsen/logrus"
)

type apiOutLogger struct {
	LogFileWrite *logrus.Logger
	LogOutPut    *logrus.Logger
}

// NewOutLogger New Creates a new logger with a "stderr" writer to send
// log messages at or above lvl to standard output.
func NewOutLogger() Logger4 {
	var log = logrus.New()

	log.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true, DisableTimestamp: false, DisableColors: false, ForceColors: true, DisableSorting: false}
	log.SetLevel(logrus.DebugLevel)
	log.Out = os.Stdout
	// log.SetFormatter(logFormat)

	return &apiOutLogger{
		LogFileWrite: nil,
		LogOutPut:    log,
	}
}

func (e *apiOutLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return e.LogOutPut.WithFields(fields)
}

func (e *apiOutLogger) WithError(err error) *logrus.Entry {
	return e.LogOutPut.WithError(err)
}

func (e *apiOutLogger) Fatal(args ...interface{}) {
	e.LogOutPut.Fatal(args)
}
func (e *apiOutLogger) Fatalf(format string, args ...interface{}) {
	e.LogFileWrite.Fatalf(format, args)
}

func (e *apiOutLogger) Debug(args ...interface{}) {
	e.LogOutPut.Debug(args)
}
func (e *apiOutLogger) Debugf(format string, args ...interface{}) {
	e.LogOutPut.Debugf(format, args)
}

func (e *apiOutLogger) Warning(args ...interface{}) {
	e.LogOutPut.Warning(args)
}
func (e *apiOutLogger) Warningf(format string, args ...interface{}) {
	e.LogOutPut.Warningf(format, args)
}

func (e *apiOutLogger) Info(args ...interface{}) {
	e.LogOutPut.Info(args)
}
func (e *apiOutLogger) Infof(format string, args ...interface{}) {
	e.LogOutPut.Infof(format, args)
}

func (e *apiOutLogger) Error(args ...interface{}) {
	e.LogOutPut.Error(args)
}
func (e *apiOutLogger) Errorf(format string, args ...interface{}) {
	e.LogOutPut.Errorf(format, args)
}
