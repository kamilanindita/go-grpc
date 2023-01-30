package service

import (
	"github.com/kamilanindita/go-grpc/server/model"
	"github.com/kamilanindita/go-grpc/server/repository"
)

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{bookRepository: bookRepository}
}

func (e *BookServiceImpl) CreateBook(book *model.CreateBookRequest) (*model.BookDB, error) {
	return e.bookRepository.CreateBook(book)
}

func (e *BookServiceImpl) UpdateBook(id int, book *model.UpdateBookRequest) (*model.BookDB, error) {
	return e.bookRepository.UpdateBook(id, book)
}

func (e *BookServiceImpl) FindBookById(id int) (*model.BookDB, error) {
	return e.bookRepository.FindBookById(id)
}

func (e *BookServiceImpl) FindBooks(page int, limit int) (model.Pagination, error) {
	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	return e.bookRepository.FindBooks(page, limit)
}

func (e *BookServiceImpl) DeleteBook(id int) error {
	return e.bookRepository.DeleteBook(id)
}
