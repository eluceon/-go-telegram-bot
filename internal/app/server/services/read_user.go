package services

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/repository"
	pb "github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *server) ReadUser(ctx context.Context, req *pb.ID) (*pb.User, error) {
	user, err := s.repo.ReadUser(ctx, req.ID)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.User{
		ID:             user.ID,
		Username:       user.Username,
		CorrectAnswers: int32(user.CorrectAnswers),
		TotalAnswers:   int32(user.TotalAnswers),
		IsPassing:      user.IsPassing,
		RegisteredAt:   user.RegisteredAt.Time.Format(time.RFC3339),
	}, err
}
