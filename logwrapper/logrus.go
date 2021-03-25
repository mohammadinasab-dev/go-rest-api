package logwrapper

import (
	"fmt"
	"go-rest-api/configuration"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//STDLog is as standard log struct type
var STDLog *StandardLog

// StandardLog enforces specifics to developer
type StandardLog struct {
	*logrus.Logger
}

// NewStandardLogger initializes the standard logger
func NewStandardTextLogger(out io.Writer) *StandardLog {
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

// NewStandardLogger initializes the standard logger
func NewStandardJsonLogger(out io.Writer) *StandardLog {
	baseLogger := logrus.New()
	baseLogger.SetReportCaller(true)
	baseLogger.SetOutput(out)
	baseLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuratiom
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

func LogSetUp() error {
	err := configuration.LoadSetup(".")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := LogSetUp()
	if err != nil {
		log.Fatalln(err)
	}

	if viper.GetString("log.logout") == "file" {
		f, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println(err) //handle
		}
		if viper.GetString("log.logformat") == "json" {
			STDLog = NewStandardJsonLogger(f)
		} else {
			STDLog = NewStandardTextLogger(f)
		}
	} else {
		if viper.GetString("log.logformat") == "json" {
			STDLog = NewStandardJsonLogger(os.Stdout)
		} else {
			STDLog = NewStandardTextLogger(os.Stdout)
		}
	}
}
