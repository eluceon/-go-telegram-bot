package services

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/repository"
	"github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *server) GetTopUsers(ctx context.Context, req *userpb.Empty) (*userpb.Users, error) {
	topUsers, err := s.repo.GetTopUsers(ctx)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	var users []*userpb.User
	for _, user := range topUsers {
		users = append(users, &userpb.User{
			ID:             user.ID,
			Username:       user.Username,
			CorrectAnswers: int32(user.CorrectAnswers),
			TotalAnswers:   int32(user.TotalAnswers),
			RegisteredAt:   user.RegisteredAt.Time.Format(time.RFC3339),
		})
	}

	return &userpb.Users{Users: users}, err
}
