package restfullapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api/data"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {

	tt := []struct {
		name       string
		inputJSON  string
		statusCode int
		response   string
	}{
		{
			name:       "first_book",
			inputJSON:  `{"title":"firstTestBook", "gener":"test","description":"first test book at srtory book"}`,
			statusCode: 200,
			response:   "start newbook",
		},
		{
			name:       "third_book",
			inputJSON:  `{"title":"thirdTestBook", "gener":"test","description":"first test book at srtory book"}`,
			statusCode: 400,
			response:   "start book",
		},
		{
			name:       "fourth_book",
			inputJSON:  `{"title":"fourthTestBook", "gener":"test","description":"fourth test book at srtory book"}`,
			statusCode: 401,
			response:   "start book",
		},
	}
	err := refreshTables()
	if err != nil {
		log.Fatalln(err)
	}
	//defer refreshTables()
	token := signUp()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "localhost:8000/book", bytes.NewBufferString(tc.inputJSON))
			if err != nil {
				t.Errorf("this is the error: %v\n", err)
			}
			req.Header.Set("Authorization", token)
			rr := httptest.NewRecorder()
			h := rootHandler(Handler.newBook)
			h.ServeHTTP(rr, req)
			var responseMap string
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				fmt.Printf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, tc.statusCode, rr.Code)
			if tc.statusCode == 200 {
				assert.Equal(t, tc.response, responseMap)
			} else {
				assert.Equal(t, tc.response, responseMap)
			}
		})
	}
}

func TestGetBooks(t *testing.T) {
	err := refreshTables()
	if err != nil {
		log.Fatalln(err)
	}
	//defer refreshTables()
	_, err = addBooks()
	if err != nil {
		log.Fatal(err)
	}
	token := signUp()
	req, err := http.NewRequest("GET", "localhost:8000/book", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	req.Header.Set("Authorization", token)
	rr := httptest.NewRecorder()
	h := rootHandler(Handler.getAllBook)
	h.ServeHTTP(rr, req)
	var books []data.Book
	err = json.Unmarshal([]byte(rr.Body.String()), &books)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(books), 3)
}
