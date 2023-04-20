package service

import (
	"chall2/desc"
	"fmt"
)

// BookService is a service for managing books.

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) GetBook(bookId string) desc.Book {
	var bookDesc desc.Book

	for _, book := range desc.BookData {
		if book.BookID == bookId {
			bookDesc = book
		}
	}

	return bookDesc
}

func (s *BookService) GetBooks() []desc.Book {
	return desc.BookData
}

func (s *BookService) AddBook(request desc.Book) desc.Book {

	var bookId = fmt.Sprintf("c%d", len(desc.BookData)+1)
	request.BookID = bookId
	desc.BookData = append(desc.BookData, request)

	return request
}

func (s *BookService) UpdateBook(bookId string, request desc.Book) desc.Book {
	var updateBook desc.Book

	for index, book := range desc.BookData {
		if book.BookID == bookId {
			request.BookID = bookId
			desc.BookData[index] = request
			updateBook = desc.BookData[index]
		}
	}

	return updateBook
}

func (s *BookService) DeleteBook() bool {
	return true
}
