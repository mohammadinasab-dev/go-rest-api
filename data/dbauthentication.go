package data

import (
	"github.com/twinj/uuid"
)

func (handler *SQLHandler) FetchAuth(auth *Authentication) (*Authentication, error) {
	au := &Authentication{}
	err := handler.DB.Debug().Where("user_id = ? AND auth_uuid = ?", auth.UserID, auth.AuthUUID).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

//Once a user row in the auth table
func (handler *SQLHandler) DeleteAuth(auth *Authentication) error {
	au := &Authentication{}
	db := handler.DB.Debug().Where("user_id = ? AND auth_uuid = ?", auth.UserID, auth.AuthUUID).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

//Once the user signup/login, create a row in the auth table, with a new uuid
func (handler *SQLHandler) CreateAuth(userID int) (*Authentication, error) {
	au := &Authentication{}
	au.AuthUUID = uuid.NewV4().String() //generate a new UUID each time
	au.UserID = userID
	err := handler.DB.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}
