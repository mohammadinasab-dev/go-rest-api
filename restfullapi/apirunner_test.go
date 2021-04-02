package restfullapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-rest-api/configuration"
	"go-rest-api/data"
	jwt "go-rest-api/security/authentication"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	Log "go-rest-api/logwrapper"

	_ "github.com/go-sql-driver/mysql"
)

var Handler *StoryBookRestAPIHandler

func TestMain(m *testing.M) {
	err := Database("../")
	if err != nil {
		Log.STDLog.Fatal(err)
	}
	os.Exit(m.Run())
}
func Database(path string) error {
	configTest, err := configuration.LoadConfigTest(path)
	if err != nil {
		Log.STDLog.Error(err)
		return err
	}
	db, err := data.CreateTestDBConnection(configTest)
	if err != nil {
		Log.STDLog.Error(err)
		return err
	}
	jwt.JWTSetter(configTest.JWTKey)
	Handler = NewStoryBookRestAPIHandler(db)
	Log.STDLog.Info("api will run in TEST mode")
	return nil
}

func signUp() string {
	user := data.User{
		Name:     "zari",
		Email:    "mohammadinasab.z@gmail.com",
		Password: "123456",
	}

	jsonbyte, err := json.Marshal(&user)
	req, err := http.NewRequest("POST", "localhost:8000/signup", bytes.NewBuffer(jsonbyte))
	if err != nil {
		Log.STDLog.Fatalf("this is the error: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	h := rootHandler(Handler.signup)
	h.ServeHTTP(rr, req)
	token := rr.Header().Get("Authorization")
	fmt.Println(token)
	return token
}

func refreshTables() error {
	//Handler.dbhandler.DB.Exec("SET foreign_key_checks=0")
	// handler = NewStoryBookRestAPIHandler(DB)
	err := Handler.dbhandler.DB.DropTableIfExists(&data.Context{}, &data.Book{}, &data.User{}).Error
	if err != nil {
		return err
	}
	err = Handler.dbhandler.DB.AutoMigrate(&data.User{}, &data.Book{}, &data.Context{}).Error
	if err != nil {
		return err
	}
	Log.STDLog.Info("Successfully refreshed table(s)")
	return nil
}

func addBooks() ([]data.Book, error) {
	var books = []data.Book{
		{
			Title:       "Title 1",
			Gener:       "darama",
			Description: "Hello world 1",
		},
		{
			Title:       "Title 2",
			Gener:       "fiction",
			Description: "Hello world 2",
		},
	}
	for i := range books {
		err := Handler.dbhandler.DB.Model(&data.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed bookss table: %v", err)
		}
	}
	return books, nil
}
