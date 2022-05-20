package services

import (
	"github.com/eluceon/quizwithmebot/internal/app/server/repository"
	pb "github.com/eluceon/quizwithmebot/pkg/api/userpb"
)

var lastID int64

type server struct {
	lastID int64
	repo   repository.Repository
	pb.UnimplementedUserServiceServer
}

func New(repo repository.Repository) *server {
	return &server{repo: repo}
}
