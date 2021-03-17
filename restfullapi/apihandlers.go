package restfullapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-rest-api/data"
	Err "go-rest-api/errorhandler"
	Log "go-rest-api/logwrapper"
	jwt "go-rest-api/security/authentication"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//StoryBookRestAPIHandler is a reciever type foe http handlers
type StoryBookRestAPIHandler struct {
	dbhandler *data.SQLHandler
}

//NewStoryBookRestAPIHandler make new StoryBookRestAPIHandler
func NewStoryBookRestAPIHandler(db *data.SQLHandler) *StoryBookRestAPIHandler {
	return &StoryBookRestAPIHandler{
		dbhandler: db,
	}
}

type rootHandler func(w http.ResponseWriter, r *http.Request) error

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Execute the final handler, and deal with errors
	err := h(w, r)
	if err != nil {
		if err, ok := err.(Err.HTTPError); ok {
			err.SetResponse(w, r)
			Log.STDLog.Error(err)
			return
		}
		if err, ok := err.(Err.DBError); ok {
			err.SetResponse(w, r)
			Log.STDLog.Error(err)
			return
		}
		Log.STDLog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
}

func (handler StoryBookRestAPIHandler) signup(w http.ResponseWriter, r *http.Request) error {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return &Err.ErrorReadRequestBody{Err: err}
	}
	user := data.User{}
	err = json.Unmarshal(jsonbyte, &user)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	if err := user.Validate(); err != nil {
		return &Err.ErrorValidateRequest{Err: err}
	}
	err = handler.dbhandler.DBSignUpHandler(user)
	if err != nil {
		return err
	}
	if !jwt.SetJwtToken(w, user) {
		return &Err.ErrorJWRTokenNotSet{Err: err}
	}
	json.NewEncoder(w).Encode("registeration succeed")
	return nil

}

func (handler StoryBookRestAPIHandler) login(w http.ResponseWriter, r *http.Request) error {
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return &Err.ErrorReadRequestBody{Err: err}
	}
	user := data.User{}
	err = json.Unmarshal(jsonbyte, &user)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	storeduser, err := handler.dbhandler.DBLoginHandler(user)
	if err != nil {
		return err
	}
	jsonbyte, err = json.Marshal(storeduser)
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}

	}
	if !jwt.SetJwtToken(w, user) {
		return &Err.ErrorJWRTokenNotSet{Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonbyte)
	return nil
}

func (handler StoryBookRestAPIHandler) newBook(w http.ResponseWriter, r *http.Request) error {
	user, _ := jwt.IsAuthorized(r)
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return &Err.ErrorReadRequestBody{Err: err}
	}
	book := data.Book{}
	err = json.Unmarshal(jsonbyte, &book)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}

	err = handler.dbhandler.DBInsertBook(book, user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("start newbook")
	return nil
}

func (handler StoryBookRestAPIHandler) newContext(w http.ResponseWriter, r *http.Request) error {
	user, _ := jwt.IsAuthorized(r)
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return &Err.ErrorReadRequestBody{Err: err}
	}
	context := data.Context{}
	err = json.Unmarshal(jsonbyte, &context)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	if err := context.Validate(); err != nil {
		return &Err.ErrorValidateRequest{Err: err}
	}
	err = handler.dbhandler.DBAddContext(context, user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("add new context")
	return nil
}

func (handler StoryBookRestAPIHandler) getAllBook(w http.ResponseWriter, r *http.Request) error {
	books, err := handler.dbhandler.DBGetBooks()
	if err != nil {
		fmt.Println("1010101010")
		return (err)
	}
	jsonByte, err := json.Marshal(books)
	if err != nil {
		fmt.Println("2020202020")
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	w.Header().Add("content-type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
	return nil
}

func (handler StoryBookRestAPIHandler) getBook(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		return &Err.ErrorBadRequest{Err: errors.New("bad request")}
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		return err
	}
	book, err := handler.dbhandler.DBGetBookByID(bookID)
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return nil
}

func (handler StoryBookRestAPIHandler) readBook(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		return &Err.ErrorBadRequest{Err: errors.New("bad request")}
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		return (err)
	}
	book, err := handler.dbhandler.DBReadBookByID(bookID)
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return nil
}

func (handler StoryBookRestAPIHandler) deleteBook(w http.ResponseWriter, r *http.Request) error {
	user, _ := jwt.IsAuthorized(r)
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		return &Err.ErrorBadRequest{Err: errors.New("bad request")}
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		return err
	}

	err = handler.dbhandler.DBDeleteBookByID(bookID, user)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("book deleted succesfully")
	return nil
}

func (handler StoryBookRestAPIHandler) updateBook(w http.ResponseWriter, r *http.Request) error {
	user, _ := jwt.IsAuthorized(r)
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		return &Err.ErrorBadRequest{Err: errors.New("bad request")}
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		return err
	}
	jsonbyte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return &Err.ErrorReadRequestBody{Err: err}
	}
	book := data.Book{}
	err = json.Unmarshal(jsonbyte, &book)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{}
	}
	fmt.Println(book.BookID)
	book.BookID = bookID
	err = handler.dbhandler.DBUpdateBookByID(book, user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("book update succesfully")
	return nil

}
