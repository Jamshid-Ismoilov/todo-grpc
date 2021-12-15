package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/rustagram/template-service/genproto"
	l "github.com/rustagram/template-service/pkg/logger"
	"github.com/rustagram/template-service/storage"
)

// UserService is an object that implements user interface.
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ...
func NewUserService(storage storage.IStorage, log l.Logger) *UserService {
	return &UserService{
		storage: storage,
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Create(*req)
	if err != nil {
		s.logger.Error("failed to create user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &user, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.ByIdReq) (*pb.User, error) {
	user, err := s.storage.User().Get(req.GetId())
	if err != nil {
		s.logger.Error("failed to get user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &user, nil
}

func (s *UserService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	users, count, err := s.storage.User().List(req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to list users", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list users")
	}

	return &pb.ListResp{
		Users: users,
		Count: count,
	}, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Update(*req)
	if err != nil {
		s.logger.Error("failed to update user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := s.storage.User().Delete(req.Id)
	if err != nil {
		s.logger.Error("failed to delete user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &pb.EmptyResp{}, nil
}
