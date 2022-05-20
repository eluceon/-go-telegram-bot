package services

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
	"github.com/eluceon/quizwithmebot/internal/app/server/repository"
	pb "github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) UpdateUser(ctx context.Context, req *pb.User) (*pb.Empty, error) {
	var product = models.User{
		ID:             req.ID,
		Username:       req.Username,
		CorrectAnswers: int(req.CorrectAnswers),
		TotalAnswers:   int(req.TotalAnswers),
		IsPassing:      req.IsPassing,
	}

	err := s.repo.UpdateUser(ctx, product)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &pb.Empty{}, err
}
