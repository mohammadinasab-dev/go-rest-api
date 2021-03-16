package logwrapper

import (
	"os"

	"github.com/sirupsen/logrus"
)

//ErrorLog is as STDLog struct type
var STDLog *StandardLog

//InfoLog is a STDLog struct type
// var InfoLog *STDLog

// STDLog enforces specifics to developer
type StandardLog struct {
	*logrus.Logger
}

// Event stores messages to log later, from our standard interface
// type Event struct {
// 	id      int
// 	message string
// }

// NewErrorLogger initializes the standard logger
func NewStandardLogger() *StandardLog {
	baseLogger := logrus.New()
	baseLogger.Formatter = &logrus.TextFormatter{}
	// baseLogger.SetReportCaller(true)
	baseLogger.SetOutput(os.Stdout)
	// baseLogger.SetLevel(logrus.ErrorLevel)
	return &StandardLog{baseLogger}
}

// // NewInfoLogger initialize the standard logger
// func NewInfoLogger() *STDLog {
// 	baseLogger := logrus.New()
// 	baseLogger.Formatter = &logrus.TextFormatter{}
// 	// baseLogger.SetReportCaller(true)
// 	baseLogger.SetOutput(os.Stdout)
// 	baseLogger.SetLevel(logrus.InfoLevel)
// 	return &STDLog{baseLogger}
// }

// // Declare variables to store log messages as new Events
// var (
// 	infoRequestStart = Event{21, "Info Of Starting Request: %s"}
// 	infoRequestEnd   = Event{22, "Info Of Ending Request: %s"}
// )

// // Declare variables to store log messages as new Events
// var (
// 	errorReadRequestBody = Event{1, "Error Of Reading Request Body: %s"}
// 	errorJSONMarshal     = Event{2, "Error Of Marshaling Json: %s"}
// 	errorJSONUnMarshal   = Event{3, "Error Of UnMarshaling Json: %s"}
// 	errorDatabaseResult  = Event{4, "Error Of Database Result: %s"}
// )

// // ErrorReadRequestBody is a standard error message
// func (l *STDLog) ErrorReadRequestBody(argument interface{}) {
// 	l.Errorf(errorReadRequestBody.message, argument)
// }

// // ErrorJSONMarshal is a standard error message
// func (l *STDLog) ErrorJSONMarshal(argument interface{}) {
// 	l.Errorf(errorJSONMarshal.message, argument)
// }

// // ErrorJSONUnMarshal is a standard error message
// func (l *STDLog) ErrorJSONUnMarshal(argument interface{}) {
// 	l.Errorf(errorJSONUnMarshal.message, argument)
// }

// // ErrorDatabaseResult is a standard error message
// func (l *STDLog) ErrorDatabaseResult(argument interface{}) {
// 	l.Errorf(errorDatabaseResult.message, argument)
// }

// // InfoRequestStart is a standard error message
// func (l *STDLog) InfoRequestStart(argument interface{}) {
// 	l.Infof(infoRequestStart.message, argument)
// }

// // InfoRequestEnd is a standard error message
// func (l *STDLog) InfoRequestEnd(argument interface{}) {
// 	l.Infof(infoRequestEnd.message, argument)
// }

func init() {
	STDLog = NewStandardLogger()
	// ErrorLog.Info("initiate Errorlog ...")
	// InfoLog = NewInfoLogger()
	// InfoLog.Info("initiate Infolog ...")
}
