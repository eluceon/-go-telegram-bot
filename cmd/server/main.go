package main

import (
	"context"
	"github.com/eluceon/quizwithmebot/internal/app/server/db"
	"github.com/eluceon/quizwithmebot/internal/app/server/mw"
	"github.com/eluceon/quizwithmebot/internal/app/server/repository"
	"github.com/eluceon/quizwithmebot/internal/app/server/services"
	"github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()

	adp, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	newServer := services.New(repository.New(adp))
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	userpb.RegisterUserServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
