package main

import (
	"github.com/eluceon/quizwithmebot/internal/app/parser/app"
	"github.com/eluceon/quizwithmebot/internal/app/parser/mw"
	pb "github.com/eluceon/quizwithmebot/pkg/api/quizpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterQuizServiceServer(grpcServer, app.New())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
