package restfullapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-rest-api/data"
	Err "go-rest-api/errorhandler"
	Log "go-rest-api/logwrapper"
	"go-rest-api/response"
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
	auth, err := handler.dbhandler.DBSignUpHandler(user)
	if err != nil {
		return err
	}
	if !jwt.SetJwtToken(w, auth) {
		return &Err.ErrorJWRTokenNotSet{Err: err}
	}
	err = response.JSON(w, "true", "registeration succeed", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
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
	auth, err := handler.dbhandler.DBLoginHandler(user)
	if err != nil {
		return err
	}
	if !jwt.SetJwtToken(w, auth) {
		return &Err.ErrorJWRTokenNotSet{Err: err}
	}
	err = response.JSON(w, "true", "login succeed", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
	return nil
}

func (handler StoryBookRestAPIHandler) logout(w http.ResponseWriter, r *http.Request) error {
	auth, _ := jwt.IsValid(r)
	if err := handler.dbhandler.DeleteAuth(&auth); err != nil {
		return &Err.ErrorDBDeleteResult{Err: err}
	}
	err := response.JSON(w, "true", "log out successfully", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
	return nil
}

func (handler StoryBookRestAPIHandler) newBook(w http.ResponseWriter, r *http.Request) error {
	auth, _ := jwt.IsValid(r)
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

	err = handler.dbhandler.DBInsertBook(book, auth.UserID)
	if err != nil {
		return err
	}
	err = response.JSON(w, "true", "new book add successfully", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
	return nil
}

func (handler StoryBookRestAPIHandler) newContext(w http.ResponseWriter, r *http.Request) error {
	auth, _ := jwt.IsValid(r)
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
	err = handler.dbhandler.DBAddContext(context, auth.UserID)
	if err != nil {
		return err
	}
	err = response.JSON(w, "true", "new context add successfully", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONMarshal{Err: err}
	}
	return nil
}

func (handler StoryBookRestAPIHandler) getAllBook(w http.ResponseWriter, r *http.Request) error {
	books, err := handler.dbhandler.DBGetBooks()
	if err != nil {
		return (err)
	}
	err = response.JSON(w, "true", "books load successfully", http.StatusOK, books)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
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
	err = response.JSON(w, "true", "the book load successfully", http.StatusOK, book)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
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
	err = response.JSON(w, "true", "books load completely successfull", http.StatusOK, book)
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	return nil
}

func (handler StoryBookRestAPIHandler) deleteBook(w http.ResponseWriter, r *http.Request) error {
	auth, _ := jwt.IsValid(r)
	vars := mux.Vars(r)
	if _, ok := vars["ID"]; !ok {
		return &Err.ErrorBadRequest{Err: errors.New("bad request")}
	}
	bookID, err := strconv.Atoi(vars["ID"])
	if err != nil {
		return err
	}

	err = handler.dbhandler.DBDeleteBookByID(bookID, auth.UserID)
	if err != nil {
		return err
	}
	err = response.JSON(w, "true", "book deleted successfully", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	return nil
}

func (handler StoryBookRestAPIHandler) updateBook(w http.ResponseWriter, r *http.Request) error {
	auth, _ := jwt.IsValid(r)
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
	err = handler.dbhandler.DBUpdateBookByID(book, auth.UserID)
	if err != nil {
		return err
	}
	err = response.JSON(w, "true", "book update successfully", http.StatusOK, "")
	if err != nil {
		return &Err.ErrorJSONUnMarshal{Err: err}
	}
	return nil

}
