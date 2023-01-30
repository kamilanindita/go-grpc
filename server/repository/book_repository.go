package repository

import "github.com/kamilanindita/go-grpc/server/model"

type BookRepository interface {
	CreateBook(book *model.CreateBookRequest) (*model.BookDB, error)
	UpdateBook(id int, book *model.UpdateBookRequest) (*model.BookDB, error)
	FindBookById(int) (*model.BookDB, error)
	FindBooks(page int, limit int) (model.Pagination, error)
	DeleteBook(id int) error
}
