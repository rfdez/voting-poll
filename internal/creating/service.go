package creating

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
)

// Service is the interface that must be implemented by the creating service.
type Service interface {
	CreatePoll(ctx context.Context, id, title, description string) error
	CreateOption(ctx context.Context, id, title, description, pollID string) error
}

type service struct {
	pollRepository   voting.PollRepository
	optionRepository voting.OptionRepository
}

// NewService creates a new creating service.
func NewService(pollRepository voting.PollRepository, optionRepository voting.OptionRepository) Service {
	return &service{
		pollRepository:   pollRepository,
		optionRepository: optionRepository,
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

func (s *service) CreateOption(ctx context.Context, id, title, description, pollID string) error {
	pollIDVO, err := voting.NewPollID(pollID)
	if err != nil {
		return err
	}

	if _, err := s.pollRepository.Find(ctx, pollIDVO); err != nil {
		return err
	}

	option, err := voting.NewOption(id, title, description, pollID)
	if err != nil {
		return err
	}

	if err := s.optionRepository.Save(ctx, option); err != nil {
		return err
	}

	return nil
}
