package decreasing

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
)

// Service is the interface that must be implemented by the decreasing service.
type Service interface {
	DecreaseOptionVotes(ctx context.Context, optionID string) error
}

type service struct {
	optionRepository voting.OptionRepository
}

// NewService creates a new decreasing service.
func NewService(optionRepository voting.OptionRepository) Service {
	return &service{
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
