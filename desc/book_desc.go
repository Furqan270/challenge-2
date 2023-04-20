package desc

import "time"

type Book struct {
	BookID    string     `json:"book_id" gorm:"primary_key"`
	Title     string     `json:"title" gorm:"type:varchar(50);not null" `
	Author    string     `json:"author" gorm:"type:varchar(50);not null"  `
	Desc      string     `json:"desc" gorm:"type:varchar(50);not null"  `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var BookData []Book
