package errorhandler

import (
	"encoding/json"
	"fmt"
	Log "go-rest-api/logwrapper"
	"net/http"
)

// user session state interface
type HTTPError interface {
	error
	SetResponse(w http.ResponseWriter, r *http.Request)
}

// simple user unauthorized error
type ErrorReadRequestBody struct {
	// Wr  http.ResponseWriter
	// Re  *http.Request
	Err error
}

// simple user unauthorized error
type ErrorJSONMarshal struct {
	// Wr  http.ResponseWriter
	// Re  *http.Request
	Err error
}

// simple user unauthorized error
type ErrorJSONUnMarshal struct {
	// Wr  http.ResponseWriter
	// Re  *http.Request
	Err error
}

// simple user unauthorized error
type ErrorValidateRequest struct {
	// Wr  http.ResponseWriter
	// Re  *http.Request
	Err error
}

// simple user unauthorized error
type ErrorJWRTokenNotSet struct {
	// Wr  http.ResponseWriter
	// Re  *http.Request
	Err error
}

// simple user unauthorized error
type ErrorBadRequest struct {
	// Wr  http.ResponseWriter
	// Re  *http.Request
	Err error
}

// check if user is logged in
func (err *ErrorReadRequestBody) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorJSONMarshal) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorJSONUnMarshal) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorValidateRequest) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for validation error
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorJWRTokenNotSet) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for jwt token error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorBadRequest) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err) //////////////////////////////////set up logrus for jwt token error
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

// return error message
func (httpErr *ErrorReadRequestBody) Error() string {
	return fmt.Sprintf("Error to read request body:%v ", httpErr.Err)
}

// return error message
func (httpErr *ErrorJSONMarshal) Error() string {
	return fmt.Sprintf("Error to marshal the body:%v ", httpErr.Err)
}

// return error message
func (httpErr *ErrorJSONUnMarshal) Error() string {
	return fmt.Sprintf("Error to unmarshal the database result:%v ", httpErr.Err)
}

// return error message
func (httpErr *ErrorValidateRequest) Error() string {
	return fmt.Sprintf("Error of validation faild:%v ", httpErr.Err)
}

// return error message
func (httpErr *ErrorJWRTokenNotSet) Error() string {
	return fmt.Sprintf("Error of jwt token not set:%v ", httpErr.Err)
}

// return error message
func (httpErr *ErrorBadRequest) Error() string {
	return fmt.Sprintf("Error of jwt token not set:%v ", httpErr.Err)
}
