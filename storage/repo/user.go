package repo

import (
	pb "github.com/rustagram/template-service/genproto"
)

// UserStorageI ...
type UserStorageI interface {
	Create(pb.User) (pb.User, error)
	Get(id int64) (pb.User, error)
	List(page, limit int64) ([]*pb.User, int64, error)
	Update(pb.User) (pb.User, error)
	Delete(id int64) error
}
