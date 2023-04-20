package controller

import (
	"chall2/desc"
	"chall2/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (b *BookController) BookRoutes(r *gin.Engine) {
	r.POST("books", b.AddBook)
	r.GET("books", b.GetAllBooks)
	r.GET("/books/:book_id", b.GetBook)
	r.PUT("/books/:book_id", b.UpdateBook)
	r.DELETE("/books/:book_id", b.DeleteBook)
}

type BookController struct {
	service      *service.BookService
	DBConnection *gorm.DB
}

func NewBookController() *BookController {
	return &BookController{
		service: service.NewBookService(),
	}
}

func (b *BookController) AddBook(ctx *gin.Context) {
	var book desc.Book
	err := ctx.BindJSON(&book)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		//ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := b.service.AddBook(book)

	// car.CarID = fmt.Sprintf("c%d", len(model.CarData)+1)
	// model.CarData = append(model.CarData, car)

	// model.CarData = append(model.CarData, car)
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Created",
		"data":       result,
	})
}

func (b *BookController) GetAllBooks(ctx *gin.Context) {

	result := b.service.GetBooks()

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "success get all books",
		"data":       result,
	})
}

func (b *BookController) GetBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")

	for _, book := range desc.BookData {
		if book.BookID == bookID {
			ctx.JSON(200, gin.H{
				"statusCode": 200,
				"message":    "Created",
				"data":       book,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"statusCode": 404,
		"message":    "book not found",
		"data":       nil,
	})
}

func (b *BookController) UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")

	for index, book := range desc.BookData {
		if book.BookID == bookID {
			var newBook desc.Book
			err := ctx.BindJSON(&newBook)

			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			newBook.BookID = bookID

			desc.BookData[index] = newBook

			ctx.JSON(200, gin.H{
				"statusCode": 200,
				"message":    "Updated",
				"data":       newBook,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"statusCode": 404,
		"message":    "book not found",
		"data":       nil,
	})
}

func (b *BookController) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("book_id")

	for index, book := range desc.BookData {
		if book.BookID == bookID {
			desc.BookData = append(desc.BookData[:index], desc.BookData[index+1:]...)

			ctx.JSON(200, gin.H{
				"statusCode": 200,
				"message":    "Deleted",
				"data":       nil,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"statusCode": 404,
		"message":    "book not found",
		"data":       nil,
	})
}
