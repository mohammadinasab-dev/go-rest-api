package middleware

import (
	"fmt"
	Log "go-rest-api/logwrapper"
	"go-rest-api/security/jwt"
	"net/http"

	"github.com/gorilla/handlers"
)

//LoggerMiddle every request to server
func LoggerMiddle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log.InfoLog.InfoRequestStart(r.URL.Path)
		defer Log.InfoLog.InfoRequestEnd(r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

//ContentTypeMiddle content type
func ContentTypeMiddle(h http.Handler) http.Handler {
	return handlers.ContentTypeHandler(h, "application/json")
}

//JWTMiddle check jwt token
func JWTMiddle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := jwt.IsLogedin(r)
		switch ok {
		case true:
			fmt.Fprintln(w, user.Email, "you are loged in before")
			h.ServeHTTP(w, r)
		case false:
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("first you should log in"))
			}(w, r)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return

		}
	})

}

// func JWTMiddle(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user, ok := jwt.IsLogedin(r)
// 		switch ok {
// 		case true:
// 			fmt.Fprintln(w, user.Email, "you are loged in before")
// 			h(w, r)
// 		case false:
// 			func(w http.ResponseWriter, r *http.Request) {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte("first you should log in"))
// 			}(w, r)
// 		default:
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("bad request"))
// 			return

// 		}
// 	}
// }

// type middleware func(http.HandlerFunc) http.HandlerFunc

//MultipleMiddleware set middlewares by order
// func MultipleMiddleware(h http.HandlerFunc, m ...middleware) http.HandlerFunc {
// 	if len(m) < 1 {
// 		return h
// 	}
// 	wrapped := h
// 	for i := len(m) - 1; i >= 0; i-- {
// 		wrapped = m[i](wrapped)
// 	}
// 	return wrapped
// }

//CommonMiddleware set common middlewares in a slice
// var CommonMiddleware = []middleware{
// 	LoggerMiddle, JWTMiddle,
// }
