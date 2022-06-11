package increasing

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
)

// Service is the interface that must be implemented by the increasing service.
type Service interface {
	IncreasePollVoters(ctx context.Context, pollID string) error
	IncreaseOptionVotes(ctx context.Context, optionID string) error
}

type service struct {
	pollRepository   voting.PollRepository
	optionRepository voting.OptionRepository
}

// NewService creates a new increasing service.
func NewService(pollRepository voting.PollRepository, optionRepository voting.OptionRepository) Service {
	return &service{
		pollRepository:   pollRepository,
		optionRepository: optionRepository,
	}
}

func (s *service) IncreaseOptionVotes(ctx context.Context, optionID string) error {
	optionIDVO, err := voting.NewOptionID(optionID)
	if err != nil {
		return err
	}

	option, err := s.optionRepository.Find(ctx, optionIDVO)
	if err != nil {
		return err
	}

	option.IncreaseVotes()

	return s.optionRepository.Save(ctx, option)
}

func (s *service) IncreasePollVoters(ctx context.Context, pollID string) error {
	pollIDVO, err := voting.NewPollID(pollID)
	if err != nil {
		return err
	}

	poll, err := s.pollRepository.Find(ctx, pollIDVO)
	if err != nil {
		return err
	}

	poll.IncreaseVoters()

	return s.pollRepository.Save(ctx, poll)
}
