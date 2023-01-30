package service

import "github.com/kamilanindita/go-grpc/server/model"

type BookService interface {
	CreateBook(*model.CreateBookRequest) (*model.BookDB, error)
	UpdateBook(id int, book *model.UpdateBookRequest) (*model.BookDB, error)
	FindBookById(id int) (*model.BookDB, error)
	FindBooks(page int, limit int) (model.Pagination, error)
	DeleteBook(id int) error
}
