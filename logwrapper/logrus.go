package logwrapper

import (
	"os"

	"github.com/sirupsen/logrus"
)

//STDLog is as standard log struct type
var STDLog *StandardLog

// StandardLog enforces specifics to developer
type StandardLog struct {
	*logrus.Logger
}

// NewStandardLogger initializes the standard logger
func NewStandardLogger() *StandardLog {
	baseLogger := logrus.New()
	baseLogger.Formatter = &logrus.TextFormatter{}
	baseLogger.SetReportCaller(true)
	baseLogger.SetOutput(os.Stdout)
	// baseLogger.SetLevel(logrus.ErrorLevel)
	return &StandardLog{baseLogger}
}
func init() {
	STDLog = NewStandardLogger()
}
