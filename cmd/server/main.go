package main

import (
	"log"

	"github.com/Qjoyboy/users-service/internal/database"
	"github.com/Qjoyboy/users-service/internal/transport/grpc"
	"github.com/Qjoyboy/users-service/internal/user"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	repo := user.NewTaskRepository(db)
	svc := user.NewUserService(repo)

	if err := grpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой %v", err)
	}
}
