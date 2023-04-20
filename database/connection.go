package database

import (
	"chall2/desc"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
	db  *sql.DB
)

func DBConnection() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=Furq0n27* dbname=Book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection Failed to Open %v\n", err)
		panic(err)
		return db, err
	}

	db.Debug().AutoMigrate(&desc.Book{})

	fmt.Println("Connection Established")

	return db, nil

}

func AddBook() {
	var books = &desc.Book{}
	sqlStatement := `
	INSERT INTO Book (title, author, desc)
	values ($title, $author, $desc)
	returning*
	`
	err = db.QueryRow(sqlStatement, "test", "furqon", "test1").
		Scan(&books.BookID, &books.Title, &books.Author, &books.Desc)
	if err != nil {
		panic(err)
	}
	fmt.Println("sukses connected")
	AddBook()
}
func GetDB() *gorm.DB {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	return db
}
