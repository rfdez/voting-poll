package creator

import "context"

// Service is the interface that must be implemented by the creator service.
type Service interface {
	CreatePoll(ctx context.Context, id, title, description string) error
}

type service struct {
}

// NewService creates a new creator service.
func NewService() Service {
	return &service{}
}

func (s *service) CreatePoll(ctx context.Context, id, title, description string) error {
	return nil
}
