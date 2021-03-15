package data

import (
	"errors"
	Log "go-rest-api/logwrapper"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		Log.ErrorLog.Error(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

//DBSignUpHandler is sign up
func (handler *SQLHandler) DBSignUpHandler(user User) error {
	hashedPwd := hashAndSalt([]byte(user.Password))
	user.Password = hashedPwd
	result := handler.db.Debug().Create(&user)
	if result.Error != nil {
		Log.ErrorLog.Error(result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		Log.ErrorLog.Error(errors.New("no row effected"))
		return errors.New("no row effected")
	}
	return nil
}

//DBLoginHandler is log in
func (handler *SQLHandler) DBLoginHandler(user User) (User, error) {
	plainPwd := []byte(user.Password)
	stdu := User{}
	if result := handler.db.Debug().Where("email = ?", user.Email).First(&stdu); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		Log.ErrorLog.Error(result.Error)
		return User{}, result.Error
	}
	hashedPwd := stdu.Password
	if !comparePasswords(hashedPwd, plainPwd) {
		err1 := errors.New("password incorrect")
		return stdu, err1
	}
	return stdu, nil
}
