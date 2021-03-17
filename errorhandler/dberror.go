package errorhandler

import (
	"encoding/json"
	"fmt"
	Log "go-rest-api/logwrapper"
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
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBCreateResult) Error() string {
	return fmt.Sprintf("Error of database find query:%v ", httpErr.Err)
}

func (err *ErrorDBFindResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ErrorDBFindResult")
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBFindResult) Error() string {
	return fmt.Sprintf("Error of database find query:%v ", httpErr.Err)
}

func (err *ErrorDBNoRowsAffected) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (httpErr *ErrorDBNoRowsAffected) Error() string {
	return fmt.Sprintf("Error of database no row affected:%v ", httpErr.Err)
}

func (err *ErrorDBDeleteResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (httpErr *ErrorDBDeleteResult) Error() string {
	return fmt.Sprintf("Error of database no roe affected:%v ", httpErr.Err)
}

func (err *ErrorDBUpdateResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (httpErr *ErrorDBUpdateResult) Error() string {
	return fmt.Sprintf("Error of database no roe affected:%v ", httpErr.Err)
}
