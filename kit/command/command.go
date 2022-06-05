package command

import "context"

// Bus defines the expected behavior of a command bus.
type Bus interface {
	// Dispatch dispatches a command to the command bus.
	Dispatch(context.Context, Command) error
	// Register registers a command handler with the command bus.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

// Type is a type of command.
type Type string

// Command represents an application command.
type Command interface {
	// Type returns the type of command.
	Type() Type
}

// Handler defines the expected behavior of a command handler.
type Handler interface {
	// Handle handles a command.
	Handle(context.Context, Command) error
}
