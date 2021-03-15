package data

import (
	"fmt"
	"go-rest-api/configuration"
	Log "go-rest-api/logwrapper"

	"github.com/jinzhu/gorm"
)

//SQLHandler is a type
type SQLHandler struct {
	db *gorm.DB
}

// <username>:<pw>@tcp(<HOST>:<port>)/<dbname>
//CreateDBConnection is a function
func CreateDBConnection(config configuration.Config) (*SQLHandler, error) {
	connstring := fmt.Sprintf(config.DBUsername + ":" + config.DBPassword + "@" + config.DBAddress + "/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println(connstring)
	db, err := gorm.Open(config.DBDriver, connstring)
	if err != nil {
		Log.ErrorLog.Error(err)
		return nil, err
	}
	db.AutoMigrate(&User{}, &Book{}, &Context{})
	db.Model(&Book{}).AddForeignKey("user_id", "users(user_id)", "SET NULL", "CASCADE")
	db.Model(&Context{}).AddForeignKey("user_id", "users(user_id)", "SET NULL", "CASCADE")
	db.Model(&Context{}).AddForeignKey("book_id", "books(book_id)", "SET NULL", "CASCADE")
	return &SQLHandler{
		db: db,
	}, nil

}
