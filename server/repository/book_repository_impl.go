package repository

import (
	"math"

	"github.com/kamilanindita/go-grpc/server/model"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{db: db}
}

func (e *BookRepositoryImpl) CreateBook(book *model.CreateBookRequest) (*model.BookDB, error) {
	var newBook = model.BookDB{
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}
	err := e.db.Table("book").Save(&newBook).Error
	if err != nil {
		return nil, err
	}
	return &newBook, nil
}

func (e *BookRepositoryImpl) UpdateBook(id int, book *model.UpdateBookRequest) (*model.BookDB, error) {
	var updateBook model.UpdateBookRequest

	err := e.db.Table("book").Where("id = ?", id).First(&updateBook).Updates(&book).Error
	if err != nil {
		return nil, err
	}

	response := &model.BookDB{
		Id:          id,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}

	return response, nil
}

func (e *BookRepositoryImpl) FindBookById(id int) (*model.BookDB, error) {
	var book model.BookDB
	err := e.db.Table("book").Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (e *BookRepositoryImpl) FindBooks(page int, limit int) (model.Pagination, error) {
	var books []*model.BookDB
	var totalItems int64

	e.db.Model(&books).Count(&totalItems)

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	err := e.db.Scopes(NewPaginate(limit, page).PaginatedResult).Find(&books).Error
	if err != nil {
		return model.Pagination{}, err
	}

	var items []interface{}

	for _, book := range books {
		items = append(items, book)
	}

	responses := model.Pagination{
		Limit:      limit,
		Page:       page,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Items:      items,
	}

	return responses, nil
}

func (e *BookRepositoryImpl) DeleteBook(id int) error {
	var book model.BookDB
	err := e.db.Table("book").Where("id = ?", id).First(&book).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
