package errorhandler

import (
	"fmt"
	Log "go-rest-api/logwrapper"
	"go-rest-api/response"
	"net/http"
)

// Database Error Interface
type DBError interface {
	error
	SetResponse(w http.ResponseWriter, r *http.Request)
}

// simple DB create-insert query error
type ErrorDBCreateResult struct {
	Err error
}

// simple DB find-select query error
type ErrorDBFindResult struct {
	Err error
}

// simple DB delete query error
type ErrorDBDeleteResult struct {
	Err error
}

// simple DB update query error
type ErrorDBUpdateResult struct {
	Err error
}

// simple DB NoRowsAffected error
type ErrorDBNoRowsAffected struct {
	Err error
}

func (err *ErrorDBCreateResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of database create query", http.StatusInternalServerError, err.Err)
}

// return error message
func (err *ErrorDBCreateResult) Error() string {

	return fmt.Sprint(err.Err)
}

func (err *ErrorDBFindResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of database find query", http.StatusInternalServerError, err.Err)
}

// return error message
func (err *ErrorDBFindResult) Error() string {
	return fmt.Sprint(err.Err)
}

func (err *ErrorDBNoRowsAffected) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of database no row affected", http.StatusInternalServerError, err.Err)
}

func (err *ErrorDBNoRowsAffected) Error() string {
	return fmt.Sprint(err.Err)
}

func (err *ErrorDBDeleteResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of database no roe affected:", http.StatusInternalServerError, err.Err)
}

func (err *ErrorDBDeleteResult) Error() string {
	return fmt.Sprint(err.Err)
}

func (err *ErrorDBUpdateResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of database no roe affected", http.StatusInternalServerError, err.Err)
}

func (err *ErrorDBUpdateResult) Error() string {
	return fmt.Sprint(err.Err)
}
