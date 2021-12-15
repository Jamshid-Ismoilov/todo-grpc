package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/Jamshid-Ismoilov/todo-grpc/config"
	pb "github.com/Jamshid-Ismoilov/todo-grpc/genproto"
	"github.com/Jamshid-Ismoilov/todo-grpc/pkg/db"
	"github.com/Jamshid-Ismoilov/todo-grpc/pkg/logger"
	"github.com/Jamshid-Ismoilov/todo-grpc/service"
	"github.com/Jamshid-Ismoilov/todo-grpc/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgStorage := storage.NewStoragePg(connDB)

	userService := service.NewTaskService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
