package service

import (
	"context"
	"reflect"
	"testing"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Jamshid-Ismoilov/todo-grpc/genproto"
)

func TestTaskService_Create(t *testing.T) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client := pb.NewTaskServiceClient(conn)

	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "successful",
			input: pb.Task{
				Assignee: "Assignee Test",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2022-01-03T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
			want: pb.Task{
				Assignee: "Assignee Test",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2022-01-03T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create user", err)
			}
			got.Id = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client := pb.NewTaskServiceClient(conn)

	tests := []struct {
		name  string
		input pb.ByIdReq
		want  pb.Task
	}{
		{
			name: "successful",
			input: pb.ByIdReq{
				Id: "a1d258ab-efcf-481c-b2ae-cdabafb4ebba",
			},
			want: pb.Task{
				Assignee: "Assignee Test",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2022-01-03T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to get user", err)
			}
			got.Id = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_List(t *testing.T) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client := pb.NewTaskServiceClient(conn)

	tests := []struct {
		name  string
		input pb.ListReq
		want  pb.ListResp
	}{
		{
			name: "successful",
			input: pb.ListReq{
				Page: 1,
				Limit: 1,
			},
			want: pb.ListResp{ 
			Tasks : []*pb.Task{
				{
					Id: "9534a2fa-157a-4b10-a859-f38896779ace",
					Assignee: "Assignee Test",
					Title: "Title Test",
					Summary: "Summary Test",
					Deadline: "2022-01-03T17:31:00Z",
					Status: true,
					CreatedAt: "2021-12-12T17:31:00Z",
				},	
			},
			Count: 9,

	},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.List(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create user", err)
			}
			log.Println(got.Tasks)
			// got.Id = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}


func TestTaskService_Delete(t *testing.T) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client := pb.NewTaskServiceClient(conn)

	tests := []struct {
		name  string
		input pb.ByIdReq
		want  pb.EmptyResp
	}{
		{
			name: "successful",
			input: pb.ByIdReq{
				Id: "a1d258ab-efcf-481c-b2ae-cdabafb4ebba",
			},
			want: pb.EmptyResp{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Delete(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to delete user", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client := pb.NewTaskServiceClient(conn)

	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "successful",
			input: pb.Task{
				Id: "9534a2fa-157a-4b10-a859-f38896779ace",
				Assignee: "Updated",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2022-01-03T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
			want: pb.Task{
				Assignee: "Updated",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2022-01-03T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create user", err)
			}
			got.Id = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}
