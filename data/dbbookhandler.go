package data

import (
	"errors"
	Err "go-rest-api/errorhandler"

	"github.com/jinzhu/gorm"
)

//DBGetBooks return all the books store in database
func (handler *SQLHandler) DBGetBooks() ([]Book, error) {

	books := []Book{}
	result := handler.db.Debug().Find(&books)
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
	if result := handler.db.Debug().Where("book_id = ?", BookID).First(&book); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return nil, &Err.ErrorDBFindResult{Err: result.Error}
	}
	if result := handler.db.Debug().Where("book_id = ?", BookID).Find(&contexts); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return nil, &Err.ErrorDBFindResult{Err: result.Error}
	}
	return contexts, nil
}

//DBGetBookByID return just a single book correspond to BookID
func (handler *SQLHandler) DBGetBookByID(bookID int) (Book, error) {
	book := Book{}
	if result := handler.db.Where("book_id = ?", bookID).First(&book); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return Book{}, &Err.ErrorDBFindResult{Err: result.Error}
	}
	return book, nil
}

//DBGetContextsByID return just a single book correspond to BookID
// func (handler *SQLHandler) DBGetContextsByID(bookID int) ([]Context, error) {
// 	rows, err := handler.db.Query("SELECT * FROM `lovestory`.`context` WHERE book_idbook = ?", bookID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			log.Println(err)
// 			return nil, err
// 		}
// 		log.Println(err)
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	contexts := []Context{}
// 	for rows.Next() {
// 		g, err := getRowsDataContext(rows)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		contexts = append(contexts, g)
// 	}
// 	if rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return contexts, nil
// }

//DBDeleteBookByID delete a book from database
func (handler *SQLHandler) DBDeleteBookByID(bookID int, user User) error {
	stdu := User{}
	if result := handler.db.Debug().Where("email = ?", user.Email).First(&stdu); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return &Err.ErrorDBFindResult{Err: result.Error}
	}
	result := handler.db.Debug().Where("user_id = ?", stdu.UserID).Where("book_id = ?", bookID).Delete(&Book{})
	if result.Error != nil {
		return &Err.ErrorDBDeleteResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}

//DBUpdateBookByID delete a book from database
func (handler *SQLHandler) DBUpdateBookByID(book Book, user User) error {
	stdu := User{}
	if result := handler.db.Debug().Where("email = ?", user.Email).First(&stdu); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return &Err.ErrorDBFindResult{Err: result.Error}
	}
	result := handler.db.Debug().Model(&book).Where("user_id = ?", stdu.UserID).Update(&book)
	if result.Error != nil {
		return &Err.ErrorDBUpdateResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}

//DBInsertBook insert new book to database
func (handler *SQLHandler) DBInsertBook(book Book, user User) error {
	stdu := User{}
	if result := handler.db.Debug().Where("email = ?", user.Email).First(&stdu); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return &Err.ErrorDBFindResult{Err: result.Error}
	}
	book.UserID = stdu.UserID
	result := handler.db.Debug().Create(&book)
	if result.Error != nil {
		return &Err.ErrorDBCreateResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}

//DBAddContext add new context to available book
func (handler *SQLHandler) DBAddContext(context Context, user User) error {
	stdu := User{}
	if result := handler.db.Debug().Where("email = ?", user.Email).First(&stdu); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return &Err.ErrorDBFindResult{Err: result.Error}
	}
	context.UserID = stdu.UserID
	result := handler.db.Debug().Create(&context)
	if result.Error != nil {
		return &Err.ErrorDBCreateResult{Err: result.Error}
	}
	if result.RowsAffected == 0 {
		return &Err.ErrorDBNoRowsAffected{Err: errors.New("no row effected")}
	}
	return nil
}
