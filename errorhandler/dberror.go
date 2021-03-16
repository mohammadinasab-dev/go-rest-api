package errorhandler

import (
	"encoding/json"
	"fmt"
	Log "go-rest-api/logwrapper"
	"net/http"
)

// user session state interface
type DBError interface {
	error
	//	AddResReq(w http.ResponseWriter, r *http.Request)
	SetResponse(w http.ResponseWriter, r *http.Request)
}

// simple user unauthorized error
type ErrorDBCreateResult struct {
	Err error
}

// simple user unauthorized error
type ErrorDBFindResult struct {
	Err error
}

// simple user unauthorized error
type ErrorDBDeleteResult struct {
	Err error
}

// simple user unauthorized error
type ErrorDBUpdateResult struct {
	Err error
}

// simple user unauthorized error
type ErrorDBNoRowsAffected struct {
	Err error
}

func (err *ErrorDBCreateResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for DB Result error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBCreateResult) Error() string {
	return fmt.Sprintf("Error of database find query:%v ", httpErr.Err)
}

func (err *ErrorDBFindResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ErrorDBFindResult")
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for DB Result error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBFindResult) Error() string {
	return fmt.Sprintf("Error of database find query:%v ", httpErr.Err)
}

func (err *ErrorDBNoRowsAffected) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for DB Result error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBNoRowsAffected) Error() string {
	return fmt.Sprintf("Error of database no roe affected:%v ", httpErr.Err)
}

func (err *ErrorDBDeleteResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for DB Result error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBDeleteResult) Error() string {
	return fmt.Sprintf("Error of database no roe affected:%v ", httpErr.Err)
}

func (err *ErrorDBUpdateResult) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for DB Result error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorDBUpdateResult) Error() string {
	return fmt.Sprintf("Error of database no roe affected:%v ", httpErr.Err)
}
