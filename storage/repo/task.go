package repo

import (
	pb "github.com/Jamshid-Ismoilov/todo-grpc/genproto"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
	ListOverdue(page, limit int64, now string) ([]*pb.Task, int64, error)
}
