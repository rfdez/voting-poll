package inmemory

import (
	"context"

	"github.com/rfdez/voting-poll/kit/command"
)

// CommandBus is an in-memory implementation of the command bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus returns a new in-memory command bus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command bus interface.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	if handler, ok := b.handlers[cmd.Type()]; ok {
		return handler.Handle(ctx, cmd)
	}

	return nil
}

// Register implements the command bus interface.
func (b *CommandBus) Register(t command.Type, h command.Handler) {
	b.handlers[t] = h
}
