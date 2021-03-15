package restfullapi

import (
	"encoding/json"
	"fmt"
	"go-rest-api/data"
	Log "go-rest-api/logwrapper"
	"go-rest-api/security/jwt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//StoryBookRestAPIHandler is handler
type StoryBookRestAPIHandler struct {
	dbhandler *data.SQLHandler
}

//NewStoryBookRestAPIHandler make new StoryBookRestAPIHandler
func NewStoryBookRestAPIHandler(db *data.SQLHandler) *StoryBookRestAPIHandler {
	return &StoryBookRestAPIHandler{
		dbhandler: db,
	}
}

func (handler StoryBookRestAPIHandler) signup(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorReadRequestBody(err)
		return
	}
	user := data.User{}
	err = json.Unmarshal(jsonbyte, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONMarshal(err)
		return
	}
	if err := user.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.Error(err)
		return
	}
	err = handler.dbhandler.DBSignUpHandler(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorDatabaseResult(err)
		return
	}
	if !jwt.SetJwtToken(w, user) {
		fmt.Println("jwt not set")
		Log.ErrorLog.Warn("jwt not set")
	}
	w.Write([]byte("registeration succeed"))

}

func (handler StoryBookRestAPIHandler) login(w http.ResponseWriter, r *http.Request) {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorReadRequestBody(err)
		return
	}
	user := data.User{}
	err = json.Unmarshal(jsonbyte, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONUnMarshal(err)
		return
	}
	storeduser, err := handler.dbhandler.DBLoginHandler(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorDatabaseResult(err)
		return
	}
	//w.WriteHeader(http.StatusOK)
	jsonbyte, err = json.Marshal(storeduser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONMarshal(err)
		return
	}
	if !jwt.SetJwtToken(w, user) {
		Log.ErrorLog.Warn("jwt not set")
	}
	w.Write(jsonbyte)

}

func (handler StoryBookRestAPIHandler) newBook(w http.ResponseWriter, r *http.Request) {
	user, _ := jwt.IsLogedin(r)
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorReadRequestBody(err)
		return
	}
	book := data.Book{}
	err = json.Unmarshal(jsonbyte, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONUnMarshal(err)
		return
	}

	fmt.Println("i am ok")
	err = handler.dbhandler.DBInsertBook(book, user)
	if err != nil {
		Log.ErrorLog.ErrorDatabaseResult(err)
	}

	fmt.Fprint(w, "start newbook\n")
}

func (handler StoryBookRestAPIHandler) newContext(w http.ResponseWriter, r *http.Request) {
	user, _ := jwt.IsLogedin(r)
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorReadRequestBody(err)
		return
	}
	context := data.Context{}
	err = json.Unmarshal(jsonbyte, &context)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONUnMarshal(err)
		return
	}
	if err := context.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.Error(err)
		return
	}
	err = handler.dbhandler.DBAddContext(context, user)
	if err != nil {
		Log.ErrorLog.ErrorDatabaseResult(err)
	}

	fmt.Fprint(w, "add new context\n")

}

func (handler StoryBookRestAPIHandler) getAllBook(w http.ResponseWriter, r *http.Request) {
	books, err := handler.dbhandler.DBGetBooks()
	if err != nil {
		Log.ErrorLog.ErrorDatabaseResult(err)
	}
	jsonByte, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONMarshal(err)
	}
	w.Header().Add("content-type", "aplication/json") //Header().Add???**********************
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)

}

func (handler StoryBookRestAPIHandler) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		Log.ErrorLog.Infof("Bad Url Request: %s", r.URL.String())
		return
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.Error(err)
	}
	book, err := handler.dbhandler.DBGetBookByID(bookID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no row founded!"))
		Log.ErrorLog.ErrorDatabaseResult(err)
	} else {
		jsonBytes, err := json.Marshal(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			Log.ErrorLog.ErrorJSONMarshal(err)
		} else {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)
		}
	}
}

func (handler StoryBookRestAPIHandler) readBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		Log.ErrorLog.Infof("Bad Url Request: %s", r.URL.String())
		return
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.Error(err)
	}
	tbook, err := handler.dbhandler.DBReadBookByID(bookID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorDatabaseResult(err)
	} else {
		jsonBytes, err := json.Marshal(tbook)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			Log.ErrorLog.ErrorJSONMarshal(err)
		} else {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)
		}
	}

}

func (handler StoryBookRestAPIHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	user, _ := jwt.IsLogedin(r)
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		Log.ErrorLog.Infof("Bad Url Request: %s", r.URL.String())
		return
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.Error(err)
	}

	err = handler.dbhandler.DBDeleteBookByID(bookID, user)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("book doesn't deleted"))
		Log.ErrorLog.ErrorDatabaseResult(err)
	} else {
		w.Write([]byte("book deleted succesfully"))
	}
}

func (handler StoryBookRestAPIHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	user, _ := jwt.IsLogedin(r)
	jsonbyte, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.Error(err)
		return
	}
	book := data.Book{}
	err = json.Unmarshal(jsonbyte, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorJSONUnMarshal(err)
		return
	}
	fmt.Println(book.BookID)
	err = handler.dbhandler.DBUpdateBookByID(book, user)
	if err != nil {
		w.Write([]byte(err.Error()))
		Log.ErrorLog.ErrorDatabaseResult(err)
		return
	}
	fmt.Fprint(w, "start updatebook\n")

}
