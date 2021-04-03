package data

import "time"

// User model correspond to user DB table
type User struct {
	UserID    int        `gorm:"primary_key" json:"userid"`
	Name      string     `json:"name"`
	Email     string     `gorm:"NOT NULL; UNIQUE" json:"email"`
	Password  string     `gorm:"NOT NULL; UNIQUE" json:"password"`
	CreatedAt time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW(); ON UPDATE now()" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleated_at"`
	Books     []Book     `gorm:"ForeignKey:UserID"`
	Contexts  []Context  `gorm:"ForeignKey:UserID"`
}

// Book model correspond to book DB table
type Book struct {
	BookID      int        `gorm:"primary_key" json:"bookid"`
	Title       string     `json:"title"`
	Gener       string     `json:"gener"`
	Description string     `gorm:"TYPE:TEXT" json:"description"`
	CreatedAt   time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW()" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW(); ON UPDATE now()" json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleated_at"`
	UserID      int        `json:"userid"`
	Contexts    []Context  `gorm:"ForeignKey:UserID"`
}

//Context model correspond to context DB table
type Context struct {
	ContextID int        `gorm:"primary_key" json:"contextid"`
	Title     string     `json:"title"`
	Txt       string     `gorm:"TYPE:LONGTEXT;not null" json:"text"`
	CreatedAt time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"TIMESTAMP; NOT NULL; DEFAULT NOW(); ON UPDATE now()" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleated_at"`
	UserID    int        `json:"userid"`
	BookID    int        `json:"bookid"`
}

//Auth model correspond to Auth DB table
type Auth struct {
	AuthID   uint64 `gorm:"primary_key;auto_increment" json:"authid"`
	UserID   uint64 `gorm:";NOT NULL;" json:"userid"`
	AuthUUID string `gorm:"size:255;NOT NULL;" json:"authuuid"`
}
