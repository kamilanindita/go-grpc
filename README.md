# go-grpc
Simple CRUD gRPC and REST API

## Dependencies
- github.com/gofiber/fiber/v2
- gorm.io/gorm
- gorm.io/driver/mysql  
- google.golang.org/grpc

## Setup config `.env`
Copy file `.env.example` to `.env`.

## Install Dependencies
```bash
$ go mod download
```

## Running the Service
```bash
$ go run mian.go
```

## The endpoint rest api (client service)
    Customer service: http://localhost:3000/api/book <br>
    1. GET: http://localhost:3000/api/book?page={page}&limit={limit}
    2. GET: http://localhost:3000/api/book/{id}
    3. POST: http://localhost:3000/api/book
    4. PUT: http://localhost:3000/api/book/{id}
    5. DELETE: http://localhost:3000/api/book/{id}
    
