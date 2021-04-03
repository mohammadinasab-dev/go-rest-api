package errorhandler

import (
	"fmt"
	Log "go-rest-api/logwrapper"
	"go-rest-api/response"
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
	response.ERROR(w, "false", "Error to read request body", http.StatusInternalServerError, err.Err)
}

func (httpErr *ErrorReadRequestBody) Error() string {
	return fmt.Sprint(httpErr.Err)
}

func (err *ErrorJSONMarshal) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error to marshal the body", http.StatusBadRequest, err.Err)
}

func (httpErr *ErrorJSONMarshal) Error() string {
	return fmt.Sprint(httpErr.Err)
}

func (err *ErrorJSONUnMarshal) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error to unmarshal the Json format", http.StatusBadRequest, err.Err)
}

func (httpErr *ErrorJSONUnMarshal) Error() string {
	return fmt.Sprint(httpErr.Err)
}

func (err *ErrorValidateRequest) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of validation faild", http.StatusBadRequest, err.Err)
}

func (httpErr *ErrorValidateRequest) Error() string {
	return fmt.Sprint(httpErr.Err)
}

func (err *ErrorJWRTokenNotSet) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of jwt token not set:%v ", http.StatusInternalServerError, err.Err)
}

func (httpErr *ErrorJWRTokenNotSet) Error() string {
	return fmt.Sprint(httpErr.Err)
}

func (err *ErrorBadRequest) SetResponse(w http.ResponseWriter, r *http.Request) {
	Log.STDLog.Error(err)
	response.ERROR(w, "false", "Error of jwt token not set:%v ", http.StatusInternalServerError, err.Err)
}

func (httpErr *ErrorBadRequest) Error() string {
	return fmt.Sprint(httpErr.Err)
}
