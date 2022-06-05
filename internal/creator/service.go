package creator

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
)

// Service is the interface that must be implemented by the creator service.
type Service interface {
	CreatePoll(ctx context.Context, id, title, description string) error
}

type service struct {
	pollRepository voting.PollRepository
}

// NewService creates a new creator service.
func NewService(pollRepository voting.PollRepository) Service {
	return &service{
		pollRepository: pollRepository,
	}
}

func (s *service) CreatePoll(ctx context.Context, id, title, description string) error {
	poll, err := voting.NewPoll(id, title, description)
	if err != nil {
		return err
	}

	if err := s.pollRepository.Save(ctx, poll); err != nil {
		return err
	}

	return nil
}
