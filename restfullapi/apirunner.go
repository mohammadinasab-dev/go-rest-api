package restfullapi

import (
	"fmt"
	"go-rest-api/configuration"
	"go-rest-api/data"
	Log "go-rest-api/logwrapper"
	"go-rest-api/middleware"
	jwt "go-rest-api/security/authentication"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

// var Stdlog = logwrapper.NewSTDLogger()

//RunAPI initialize the API
//Create database connection
//Set server address and port
func RunAPI(path string) error {
	Log.STDLog.Info("logrus set up")
	environment, err := configuration.LoadSetup(".")
	if err != nil {
		log.Fatalf("this %v Error was occured til loading setup file", err)
	}
	if environment == "product" {
		config, err := configuration.LoadConfig(path)
		if err != nil {
			Log.STDLog.Fatal(err)
		}
		db, err := data.CreateDBConnection(config)
		if err != nil {
			Log.STDLog.Fatal(err)
		}
		jwt.JWTSetter(config.JWTKey)
		addr := config.ServerAddress
		mux := mux.NewRouter()
		RunAPIOnRouter(mux, db)
		Log.STDLog.Info("Server Started")
		return http.ListenAndServe(addr, mux)
	}
	if environment == "test" {
		fmt.Println("TEST")
		configTest, err := configuration.LoadConfigTest(path)
		if err != nil {
			Log.STDLog.Fatal(err)
		}
		db, err := data.CreateTestDBConnection(configTest)
		if err != nil {
			Log.STDLog.Fatal(err)
		}
		jwt.JWTSetter(configTest.JWTKey)
		addr := configTest.ServerAddress
		mux := mux.NewRouter()
		RunAPIOnRouter(mux, db)
		Log.STDLog.Info("Test Server Started")
		return http.ListenAndServe(addr, mux)
	}
	return nil //CHECK
}

//RunAPIOnRouter sets the router
func RunAPIOnRouter(r *mux.Router, db *data.SQLHandler) {
	fmt.Println(runtime.Caller(1))
	handler := NewStoryBookRestAPIHandler(db)
	r.Handle("/signup", rootHandler(handler.signup)).Methods("POST")
	r.Handle("/login", rootHandler(handler.login)).Methods("POST")
	rb := r.PathPrefix("/book").Subrouter()
	rb.Handle("/", rootHandler(handler.getAllBook)).Methods("GET")
	rb.Handle("/", rootHandler(handler.newBook)).Methods("POST")
	rb.Handle("/newcontext", rootHandler(handler.newContext)).Methods("POST")
	rb.Handle("/{ID}", rootHandler(handler.getBook)).Methods("GET")
	rb.Handle("/{ID}", rootHandler(handler.updateBook)).Methods("PUT")
	rb.Handle("/{ID}", rootHandler(handler.deleteBook)).Methods("DELETE")
	rb.Handle("/view/{ID}", rootHandler(handler.getBook)).Methods("GET")
	rb.Handle("/read/{ID}", rootHandler(handler.readBook)).Methods("GET")
	r.Use(middleware.LoggerMiddle, middleware.ContentTypeMiddle)
	rb.Use(middleware.JWTMiddle)
}
