package middleware

import (
	"errors"
	Log "go-rest-api/logwrapper"
	"go-rest-api/response"
	jwt "go-rest-api/security/authentication"
	"net/http"

	"github.com/gorilla/handlers"
)

//LoggerMiddle log every incomming request to the server
func LoggerMiddle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log.STDLog.Info("api: " + r.URL.Path + " Started")
		defer Log.STDLog.Info("api: " + r.URL.Path + " Ended")
		h.ServeHTTP(w, r)
	})
}

//ContentTypeMiddle checks content type of incomming request to the server
func ContentTypeMiddle(h http.Handler) http.Handler {
	return handlers.ContentTypeHandler(h, "application/json")
}

//JWTMiddle checks jwt token
func JWTMiddle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := jwt.IsAuthorized(r)
		switch ok {
		case true:
			h.ServeHTTP(w, r)
		case false:
			func(w http.ResponseWriter, r *http.Request) {
				response.ERROR(w, "false", "Error of Unauthorized user", http.StatusUnauthorized, errors.New("Unauthorized attempt"))
			}(w, r)
		default:
			response.ERROR(w, "false", "Error of Unauthorized user", http.StatusBadRequest, errors.New("BadRequest"))
		}
	})

}
