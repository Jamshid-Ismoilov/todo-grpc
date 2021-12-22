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

var ID, UpdatedAT string


func TestTaskService_Create(t *testing.T) {

	// conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did not connect %v", err)
	// }
	// client := pb.NewTaskServiceClient(conn)

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
				Deadline: "2021-01-01T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
			want: pb.Task{
				Assignee: "Assignee Test",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2021-01-01T17:31:00Z",
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
			ID = got.Id
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
				Id: ID,
			},
			want: pb.Task{
				Assignee: "Assignee Test",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2021-01-01T17:31:00Z",
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
					Id: ID,
					Assignee: "Assignee Test",
					Title: "Title Test",
					Summary: "Summary Test",
					Deadline: "2021-01-01T17:31:00Z",
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
			
			//this line used becouse get test always adds one count so tester should change tc.want.Count everytime 
			tc.want.Count = got.Count
			
			// got.Id = ""
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
				Id: ID,
				Assignee: "Updated",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2023-01-01T17:31:00Z",
				Status: true,
				CreatedAt: "2021-12-12T17:31:00Z",
			},
			want: pb.Task{
				Assignee: "Updated",
				Title: "Title Test",
				Summary: "Summary Test",
				Deadline: "2023-01-01T17:31:00Z",
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

			//can't get the exact time of updating to "want" so copied from result
			tc.want.UpdatedAt = got.UpdatedAt
			UpdatedAT = got.UpdatedAt
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_ListOverdue(t *testing.T) {

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
					Id: ID,
					Assignee: "Updated",
					Title: "Title Test",
					Summary: "Summary Test",
					Deadline: "2023-01-01T17:31:00Z",
					Status: true,
					CreatedAt: "2021-12-12T17:31:00Z",
					UpdatedAt: UpdatedAT,
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
			
			//this line used becouse get test always adds one count so tester should change tc.want.Count everytime 
			tc.want.Count = got.Count
			
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
				Id: ID,
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
