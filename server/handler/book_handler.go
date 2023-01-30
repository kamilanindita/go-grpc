package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kamilanindita/go-grpc/server/model"
	"github.com/kamilanindita/go-grpc/server/pb"
	"github.com/kamilanindita/go-grpc/server/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookHandler struct {
	pb.UnimplementedBookServiceServer
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) (BookHandler, error) {
	bookHandler := BookHandler{
		bookService: bookService,
	}

	return bookHandler, nil
}

func (bookHandler *BookHandler) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {

	newBook := &model.CreateBookRequest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Author:      req.GetAuthor(),
	}
	createNewBook, err := bookHandler.bookService.CreateBook(newBook)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbBook := &pb.Book{
		Id:          int64(createNewBook.Id),
		Title:       createNewBook.Title,
		Description: createNewBook.Description,
		Author:      createNewBook.Author,
	}

	book := &pb.BookResponse{
		Book: pbBook,
	}

	return book, nil
}

func (bookHandler *BookHandler) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.BookResponse, error) {
	id := req.GetId()

	updateBook := model.UpdateBookRequest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Author:      req.GetAuthor(),
	}
	updatedBook, err := bookHandler.bookService.UpdateBook(int(id), &updateBook)

	fmt.Print(updateBook)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbBook := &pb.Book{
		Id:          int64(updatedBook.Id),
		Title:       updatedBook.Title,
		Description: updatedBook.Description,
		Author:      updatedBook.Author,
	}

	book := &pb.BookResponse{
		Book: pbBook,
	}

	return book, nil

}

func (bookHandler *BookHandler) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.BookResponses, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()

	books, err := bookHandler.bookService.FindBooks(int(page), int(limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var items []*pb.Book

	// Convert map to json string
	jsonStr, _ := json.Marshal(books.Items)
	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, &items); err != nil {
		fmt.Println(err)
	}
	responses := &pb.BookResponses{
		Limit:      int64(books.Limit),
		Page:       int64(books.Page),
		TotalItems: int64(books.TotalItems),
		TotalPages: int64(books.TotalPages),
		Items:      items,
	}

	return responses, nil
}

func (bookHandler *BookHandler) GetBook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	id := req.GetId()
	findBook, err := bookHandler.bookService.FindBookById(int(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbBook := &pb.Book{
		Id:          int64(findBook.Id),
		Title:       findBook.Title,
		Description: findBook.Description,
		Author:      findBook.Author,
	}

	book := &pb.BookResponse{
		Book: pbBook,
	}

	return book, nil
}

func (bookHandler *BookHandler) DeleteBook(ctx context.Context, req *pb.BookRequest) (*pb.DeleteBookResponse, error) {
	id := req.GetId()
	err := bookHandler.bookService.DeleteBook(int(id))

	var response pb.DeleteBookResponse
	if err != nil {
		response.Success = false
		return &response, status.Errorf(codes.Internal, err.Error())
	}

	response.Success = true

	return &response, nil
}
