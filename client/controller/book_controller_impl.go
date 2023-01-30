package controller

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kamilanindita/go-grpc/client/exception"
	"github.com/kamilanindita/go-grpc/client/model"
	"github.com/kamilanindita/go-grpc/client/pb"
)

type BookHandler struct {
	bookClient pb.BookServiceClient
}

// func NewNBookController(client *pb.NewBookServiceClient) BookController {
// 	return &BookControllerImpl{client: &client}
// }

func BookServiceRoute(app *fiber.App, bookClient pb.BookServiceClient) {
	bookHandler := &BookHandler{bookClient}

	app.Post("/api/book", bookHandler.Create)
	app.Put("/api/book/:id", bookHandler.Update)
	app.Get("/api/book", bookHandler.List)
	app.Get("/api/book/:id", bookHandler.FindOneById)
	app.Delete("/api/book/:id", bookHandler.Delete)
}

func (bookHandler *BookHandler) Create(c *fiber.Ctx) error {
	var request pb.CreateBookRequest

	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	createBook := &pb.CreateBookRequest{
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
	}
	responses, _ := bookHandler.bookClient.CreateBook(context.Background(), createBook)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses.GetBook(),
	})
}

func (bookHandler *BookHandler) Update(c *fiber.Ctx) error {
	var request pb.CreateBookRequest

	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	Id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	updateBook := &pb.UpdateBookRequest{
		Id:          Id,
		Title:       request.Title,
		Description: request.Description,
		Author:      request.Author,
	}
	responses, _ := bookHandler.bookClient.UpdateBook(context.Background(), updateBook)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses.GetBook(),
	})
}

func (bookHandler *BookHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	request := &pb.GetBooksRequest{
		Limit: &limit,
		Page:  &page,
	}
	responses, _ := bookHandler.bookClient.GetBooks(context.Background(), request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (bookHandler *BookHandler) FindOneById(c *fiber.Ctx) error {
	Id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	pbBookRequest := &pb.BookRequest{
		Id: Id,
	}
	responses, _ := bookHandler.bookClient.GetBook(context.Background(), pbBookRequest)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses.GetBook(),
	})
}

func (bookHandler *BookHandler) Delete(c *fiber.Ctx) error {
	Id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	pbBookRequest := &pb.BookRequest{
		Id: Id,
	}
	responses, _ := bookHandler.bookClient.DeleteBook(context.Background(), pbBookRequest)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
