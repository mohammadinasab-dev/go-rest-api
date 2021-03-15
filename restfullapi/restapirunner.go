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
		Log.ErrorLog.Fatal(err)
	}
	db, err := data.CreateDBConnection(config)
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	addr := config.ServerAddress
	mux := mux.NewRouter()
	RunAPIOnRouter(mux, db)
	Log.InfoLog.Info("Server Started ...")
	return http.ListenAndServe(addr, mux)
}

//RunAPIOnRouter is runapionrouter
func RunAPIOnRouter(r *mux.Router, db *data.SQLHandler) {
	handler := NewStoryBookRestAPIHandler(db)
	r.HandleFunc("/signup", handler.signup).Methods("POST")
	r.HandleFunc("/login", handler.login).Methods("POST")
	rb := r.PathPrefix("/book").Subrouter()
	rb.HandleFunc("/new", handler.newBook).Methods("POST")
	rb.HandleFunc("/newcontext", handler.newContext).Methods("POST")
	rb.HandleFunc("/{ID}", handler.getBook).Methods("GET")
	rb.HandleFunc("/{ID}", handler.updateBook).Methods("PUT")
	rb.HandleFunc("/{ID}", handler.deleteBook).Methods("DELETE")
	rb.HandleFunc("/view/{ID}", handler.getBook).Methods("GET")
	rb.HandleFunc("/read/{ID}", handler.readBook).Methods("GET")
	rb.HandleFunc("/all", handler.getAllBook).Methods("GET")
	// ru := r.PathPrefix("/user/").Subrouter()
	r.Use(middleware.LoggerMiddle, middleware.ContentTypeMiddle)
	rb.Use(middleware.JWTMiddle)

}
