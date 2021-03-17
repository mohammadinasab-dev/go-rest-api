package errorhandler

import (
	"encoding/json"
	"fmt"
	Log "go-rest-api/logwrapper"
	"net/http"
)

// HTTP Error Interface
type HTTPError interface {
	error
	SetResponse(w http.ResponseWriter, r *http.Request)
}

// simple http read request error
type ErrorReadRequestBody struct {
	Err error
}

// simple json marshal error
type ErrorJSONMarshal struct {
	Err error
}

// simple json marshal error
type ErrorJSONUnMarshal struct {
	Err error
}

// simple http request validation error
type ErrorValidateRequest struct {
	Err error
}

// simple http-jwt token not set error
type ErrorJWRTokenNotSet struct {
	Err error
}

// simple http bad request error
type ErrorBadRequest struct {
	Err error
}

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
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorJWRTokenNotSet) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (err *ErrorBadRequest) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func (httpErr *ErrorReadRequestBody) Error() string {
	return fmt.Sprintf("Error to read request body:%v ", httpErr.Err)
}

func (httpErr *ErrorJSONMarshal) Error() string {
	return fmt.Sprintf("Error to marshal the body:%v ", httpErr.Err)
}

func (httpErr *ErrorJSONUnMarshal) Error() string {
	return fmt.Sprintf("Error to unmarshal the database result:%v ", httpErr.Err)
}

func (httpErr *ErrorValidateRequest) Error() string {
	return fmt.Sprintf("Error of validation faild:%v ", httpErr.Err)
}

func (httpErr *ErrorJWRTokenNotSet) Error() string {
	return fmt.Sprintf("Error of jwt token not set:%v ", httpErr.Err)
}

func (httpErr *ErrorBadRequest) Error() string {
	return fmt.Sprintf("Error of jwt token not set:%v ", httpErr.Err)
}
