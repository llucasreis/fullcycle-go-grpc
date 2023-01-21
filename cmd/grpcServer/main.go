package main

import (
	"database/sql"
	"net"

	"github.com/llucasreis/fullcycle-go-grpc/internal/database"
	"github.com/llucasreis/fullcycle-go-grpc/internal/pb"
	"github.com/llucasreis/fullcycle-go-grpc/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	logrus.Infof("Created database...")
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcSever := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcSever, categoryService)

	logrus.Info("Register service...")

	reflection.Register(grpcSever)

	logrus.Info("Applied reflection...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	logrus.Info("Started listener...")
	if err := grpcSever.Serve(lis); err != nil {
		panic(err)
	}
	logrus.Infof("Running code")
}
