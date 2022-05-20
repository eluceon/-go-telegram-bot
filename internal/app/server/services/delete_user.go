package services

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/repository"
	pb "github.com/eluceon/quizwithmebot/pkg/api/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Empty, error) {
	err := s.repo.DeleteUser(ctx, req.ID)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.Empty{}, err
}
