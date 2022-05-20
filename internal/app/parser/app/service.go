package app

import (
	pb "github.com/eluceon/quizwithmebot/pkg/api/quizpb"
)

type server struct {
	pb.UnimplementedQuizServiceServer
}

func New() *server {
	return &server{}
}
