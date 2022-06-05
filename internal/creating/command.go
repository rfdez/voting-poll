package creating

import (
	"context"

	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/kit/command"
)

const (
	PollCommandType   command.Type = "voting-app.voting-poll.1.command.poll.create"
	OptionCommandType command.Type = "voting-app.voting-poll.1.command.option.create"
)

// PollCommand is a command to create a poll.
type PollCommand struct {
	id          string
	title       string
	description string
}

// NewPollCommand creates a new PollCommand.
func NewPollCommand(id, title, description string) PollCommand {
	return PollCommand{
		id:          id,
		title:       title,
		description: description,
	}
}

func (c PollCommand) Type() command.Type {
	return PollCommandType
}

// PollCommandHandler is a handler for PollCommand.
type PollCommandHandler struct {
	service Service
}

// NewPollCommandHandler creates a new PollCommandHandler.
func NewPollCommandHandler(service Service) PollCommandHandler {
	return PollCommandHandler{
		service: service,
	}
}

// Handle implements command.Handler.
func (h PollCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createPollCmd, ok := cmd.(PollCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreatePoll(
		ctx,
		createPollCmd.id,
		createPollCmd.title,
		createPollCmd.description,
	)
}

// OptionCommand is a command to create a option.
type OptionCommand struct {
	id          string
	title       string
	description string
	pollID      string
}

// NewOptionCommand creates a new OptionCommand.
func NewOptionCommand(id, title, description, pollID string) OptionCommand {
	return OptionCommand{
		id:          id,
		title:       title,
		description: description,
		pollID:      pollID,
	}
}

func (c OptionCommand) Type() command.Type {
	return OptionCommandType
}

// OptionCommandHandler is a handler for OptionCommand.
type OptionCommandHandler struct {
	service Service
}

// NewOptionCommandHandler creates a new OptionCommandHandler.
func NewOptionCommandHandler(service Service) OptionCommandHandler {
	return OptionCommandHandler{
		service: service,
	}
}

// Handle implements command.Handler.
func (h OptionCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createOptionCmd, ok := cmd.(OptionCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateOption(
		ctx,
		createOptionCmd.id,
		createOptionCmd.title,
		createOptionCmd.description,
		createOptionCmd.pollID,
	)
}
