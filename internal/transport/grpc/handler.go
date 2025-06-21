package grpc

import (
	"context"

	userpb "github.com/Qjoyboy/project-proto/proto/user"
	"github.com/Qjoyboy/users-service/internal/user"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	Svc user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.UserService) *Handler {
	return &Handler{Svc: svc}
}

func (s *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := s.Svc.CreateUser(req.Email, "")
	if err != nil {
		return nil, errors.Errorf("failed to create user: %v", err)
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    user.ID,
			Email: user.Email,
		},
	}, nil
}

func (s *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	user, err := s.Svc.GetUserByID(req.Id)
	if err != nil {
		return nil, errors.Errorf("user not found %v", err)
	}

	return &userpb.User{
		Id:    req.Id,
		Email: user.Email,
	}, nil
}

func (s *Handler) ListUsers(ctx context.Context, _ *emptypb.Empty) (*userpb.ListUsersResponse, error) {
	users, err := s.Svc.GetUsers()

	if err != nil {
		return nil, errors.Errorf("failed to list users %v", err)
	}

	var pbUsers []*userpb.User
	for _, u := range users {

		pbUsers = append(pbUsers, &userpb.User{
			Id:    u.ID,
			Email: u.Email,
		})
	}
	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (s *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	updatedUser := user.User{
		Email: req.User.Email,
	}

	usResult, err := s.Svc.UpdateUser(req.User.Id, updatedUser.Email, "")
	if err != nil {
		return nil, errors.Errorf("Failed to update user %v", err)
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    req.User.Id,
			Email: usResult.Email,
		},
	}, nil
}

func (s *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := s.Svc.DeleteUser(req.Id)
	if err != nil {
		return nil, errors.Errorf("Failed to delete user %v", err)
	}
	return &userpb.DeleteUserResponse{}, nil
}
