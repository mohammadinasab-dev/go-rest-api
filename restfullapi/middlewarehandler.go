package restfullapi

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
func (handler StoryBookRestAPIHandler) JWTMiddle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, ok := jwt.IsValid(r)
		switch ok {
		case true:
			_, err := handler.dbhandler.FetchAuth(&auth)
			if err != nil {
				func(w http.ResponseWriter, r *http.Request) {
					response.ERROR(w, "false", "Error of Unauthorized user", http.StatusUnauthorized, errors.New("Unauthorized attempt"))
				}(w, r)
			} else if err == nil {
				h.ServeHTTP(w, r)
			}
		case false:
			func(w http.ResponseWriter, r *http.Request) {
				response.ERROR(w, "false", "Error of Unauthorized user", http.StatusUnauthorized, errors.New("Unauthorized attempt"))
			}(w, r)
		default:
			response.ERROR(w, "false", "Error of Unauthorized user", http.StatusBadRequest, errors.New("BadRequest"))
		}
	})

}
