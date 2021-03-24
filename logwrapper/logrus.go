package logwrapper

import (
	"fmt"
	"runtime"
	"strings"

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
	baseLogger.SetReportCaller(true)
	baseLogger.Formatter = &logrus.TextFormatter{
		ForceColors:            true,
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf(" %s", formatFuncName(f.Function)), fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	return &StandardLog{baseLogger}
}

func formatFuncName(funcname string) string {
	s := strings.Split(funcname, ".")
	return s[len(s)-1]
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func init() {
	STDLog = NewStandardLogger()
}
