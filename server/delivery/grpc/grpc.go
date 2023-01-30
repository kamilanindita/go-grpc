package grpc

import (
	"log"
	"net"

	"github.com/kamilanindita/go-grpc/server/config"
	"github.com/kamilanindita/go-grpc/server/handler"
	"github.com/kamilanindita/go-grpc/server/pb"
	"github.com/kamilanindita/go-grpc/server/repository"
	"github.com/kamilanindita/go-grpc/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func Init(configuration config.Config, db *gorm.DB) {
	grpcServerAddress := configuration.Get("GRPC_SERVER")
	grpcServerPort := configuration.Get("GRPC_PORT")

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler, _ := handler.NewBookHandler(bookService)

	grpcServer := grpc.NewServer()

	pb.RegisterBookServiceServer(grpcServer, &bookHandler)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", grpcServerAddress+":"+grpcServerPort)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
