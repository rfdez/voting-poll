package decreasing

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
)

// Service is the interface that must be implemented by the decreasing service.
type Service interface {
	DecreasePollVoters(ctx context.Context, pollID string) error
	DecreaseOptionVotes(ctx context.Context, optionID string) error
}

type service struct {
	pollRepository   voting.PollRepository
	optionRepository voting.OptionRepository
}

// NewService creates a new decreasing service.
func NewService(pollRepository voting.PollRepository, optionRepository voting.OptionRepository) Service {
	return &service{
		pollRepository:   pollRepository,
		optionRepository: optionRepository,
	}
}

func (s *service) DecreaseOptionVotes(ctx context.Context, optionID string) error {
	optionIDVO, err := voting.NewOptionID(optionID)
	if err != nil {
		return err
	}

	option, err := s.optionRepository.Find(ctx, optionIDVO)
	if err != nil {
		return err
	}

	if err := option.DecreaseVotes(); err != nil {
		return err
	}

	return s.optionRepository.Save(ctx, option)
}

func (s *service) DecreasePollVoters(ctx context.Context, pollID string) error {
	pollIDVO, err := voting.NewPollID(pollID)
	if err != nil {
		return err
	}

	poll, err := s.pollRepository.Find(ctx, pollIDVO)
	if err != nil {
		return err
	}

	if err := poll.DecreaseVoters(); err != nil {
		return err
	}

	return s.pollRepository.Save(ctx, poll)
}
