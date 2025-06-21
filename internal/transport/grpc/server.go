package grpc

import (
	"log"
	"net"

	userpb "github.com/Qjoyboy/project-proto/proto/user"
	"github.com/Qjoyboy/users-service/internal/user"
	googlegrpc "google.golang.org/grpc"
)

func RunGRPC(svc user.UserService) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	grpcServer := googlegrpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))
	log.Println("gRPC server is listening on port 50051")

	return grpcServer.Serve(lis)

}
