package data

import (
	"errors"
	Err "go-rest-api/errorhandler"

	"github.com/jinzhu/gorm"
)

//DBGetBooks return all the books store in database
func (handler *SQLHandler) DBGetBooks() ([]Book, error) {

	books := []Book{}
	result := handler.DB.Debug().Find(&books)
	if result.Error != nil {
		//Log.ErrorLog.Error(result.Error)
		return nil, &Err.ErrorDBFindResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return nil, &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return books, nil
}

//DBReadBookByID is a function to read a whole book
func (handler *SQLHandler) DBReadBookByID(BookID int) ([]Context, error) {
	contexts := []Context{}
	book := Book{}
	if result := handler.DB.Debug().Where("book_id = ?", BookID).First(&book); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return nil, &Err.ErrorDBFindResult{Err: result.Error}
	}
	if result := handler.DB.Debug().Where("book_id = ?", BookID).Find(&contexts); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return nil, &Err.ErrorDBFindResult{Err: result.Error}
	}
	return contexts, nil
}

//DBGetBookByID return just a single book correspond to BookID
func (handler *SQLHandler) DBGetBookByID(bookID int) (Book, error) {
	book := Book{}
	if result := handler.DB.Where("book_id = ?", bookID).First(&book); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return Book{}, &Err.ErrorDBFindResult{Err: result.Error}
	}
	return book, nil
}

//DBDeleteBookByID delete a book from database
func (handler *SQLHandler) DBDeleteBookByID(bookID int, userID int) error {

	result := handler.DB.Debug().Where("user_id = ?", userID).Where("book_id = ?", bookID).Delete(&Book{})
	if result.Error != nil {
		return &Err.ErrorDBDeleteResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}

//DBUpdateBookByID delete a book from database
func (handler *SQLHandler) DBUpdateBookByID(book Book, userID int) error {
	result := handler.DB.Debug().Model(&book).Where("user_id = ?", userID).Update(&book)
	if result.Error != nil {
		return &Err.ErrorDBUpdateResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}

//DBInsertBook insert new book to database
func (handler *SQLHandler) DBInsertBook(book Book, userID int) error {
	book.UserID = userID
	result := handler.DB.Debug().Create(&book)
	if result.Error != nil {
		return &Err.ErrorDBCreateResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}

//DBAddContext add new context to available book
func (handler *SQLHandler) DBAddContext(context Context, userID int) error {
	context.UserID = userID
	result := handler.DB.Debug().Create(&context)
	if result.Error != nil {
		return &Err.ErrorDBCreateResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}
