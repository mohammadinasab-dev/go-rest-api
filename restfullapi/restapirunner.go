package restfullapi

import (
	"go-rest-api/configuration"
	"go-rest-api/data"
	Log "go-rest-api/logwrapper"
	"go-rest-api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// var Stdlog = logwrapper.NewSTDLogger()

//RunAPI initial API
//make database connection
//set server address and port
func RunAPI(filename string) error {
	config, err := configuration.LoadConfig(".")
	if err != nil {
		Log.STDLog.Fatal(err)
	}
	db, err := data.CreateDBConnection(config)
	if err != nil {
		Log.STDLog.Fatal(err)
	}
	addr := config.ServerAddress
	mux := mux.NewRouter()
	RunAPIOnRouter(mux, db)
	Log.STDLog.Info("Server Started ...")
	return http.ListenAndServe(addr, mux)
}

//RunAPIOnRouter is runapionrouter
func RunAPIOnRouter(r *mux.Router, db *data.SQLHandler) {
	handler := NewStoryBookRestAPIHandler(db)
	// r.HandleFunc("/signup", handler.signup).Methods("POST")
	r.Handle("/signup", rootHandler(handler.signup)).Methods("POST")
	r.Handle("/login", rootHandler(handler.login)).Methods("POST")
	rb := r.PathPrefix("/book").Subrouter()
	rb.Handle("/new", rootHandler(handler.newBook)).Methods("POST")
	rb.Handle("/newcontext", rootHandler(handler.newContext)).Methods("POST")
	rb.Handle("/{ID}", rootHandler(handler.getBook)).Methods("GET")
	rb.Handle("/{ID}", rootHandler(handler.updateBook)).Methods("PUT")
	rb.Handle("/{ID}", rootHandler(handler.deleteBook)).Methods("DELETE")
	rb.Handle("/view/{ID}", rootHandler(handler.getBook)).Methods("GET")
	rb.Handle("/read/{ID}", rootHandler(handler.readBook)).Methods("GET")
	rb.Handle("/all", rootHandler(handler.getAllBook)).Methods("GET")
	// ru := r.PathPrefix("/user/").Subrouter()
	r.Use(middleware.LoggerMiddle, middleware.ContentTypeMiddle)
	rb.Use(middleware.JWTMiddle)
}
