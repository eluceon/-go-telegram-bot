package services

import (
	"context"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
	pb "github.com/eluceon/quizwithmebot/pkg/api/userpb"
)

func (s *server) CreateUser(ctx context.Context, req *pb.User) (*pb.Empty, error) {
	var user = models.User{
		ID:             req.ID,
		Username:       req.Username,
		TotalAnswers:   int(req.TotalAnswers),
		CorrectAnswers: int(req.CorrectAnswers),
	}

	return &pb.Empty{}, s.repo.CreateUser(ctx, user)
}
