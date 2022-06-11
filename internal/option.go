package voting

import (
	"context"

	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/kit/uuid"
)

// OptionID represents a option identifier.
type OptionID struct {
	value string
}

// NewOptionID instantiate the VO for OptionID.
func NewOptionID(value string) (OptionID, error) {
	v, err := uuid.New(value)
	if err != nil {
		return OptionID{}, err
	}

	return OptionID{
		value: v,
	}, nil
}

// String returns the string representation of the OptionID.
func (id OptionID) String() string {
	return id.value
}

// OptionTitle represents a option title.
type OptionTitle struct {
	value string
}

// NewOptionTitle instantiate the VO for OptionTitle.
func NewOptionTitle(value string) (OptionTitle, error) {
	if value == "" {
		return OptionTitle{}, errors.NewWrongInput("poll title cannot be empty")
	}

	return OptionTitle{
		value: value,
	}, nil
}

// String returns the string representation of the OptionTitle.
func (title OptionTitle) String() string {
	return title.value
}

// OptionDescription represents a option description.
type OptionDescription struct {
	value string
}

// NewOptionDescription instantiate the VO for OptionDescription.
func NewOptionDescription(value string) OptionDescription {
	if value == "" {
		value = "No description provided"
	}

	return OptionDescription{
		value: value,
	}
}

// String returns the string representation of the OptionDescription.
func (description OptionDescription) String() string {
	return description.value
}

// OptionVotes represents the number of votes for an option.
type OptionVotes struct {
	value int
}

// NewOptionVotes creates a new OptionVotes.
func NewOptionVotes(value int) (OptionVotes, error) {
	if value < 0 {
		return OptionVotes{}, errors.NewWrongInput("option votes cannot be negative")
	}

	return OptionVotes{
		value: value,
	}, nil
}

// Value returns the Option votes.
func (votes OptionVotes) Value() int {
	return votes.value
}

// Option is the data structure that represents a poll option.
type Option struct {
	id          OptionID
	title       OptionTitle
	description OptionDescription
	pollID      PollID
	votes       OptionVotes
}

// OptionRepository is the interface that must be implemented by the Option repository.
type OptionRepository interface {
	Find(context.Context, OptionID) (Option, error)
	Save(context.Context, Option) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=OptionRepository

// NewOption creates a new Option.
func NewOption(id, name, description, pollID string, votes int) (Option, error) {
	idVO, err := NewOptionID(id)
	if err != nil {
		return Option{}, err
	}

	titleVO, err := NewOptionTitle(name)
	if err != nil {
		return Option{}, err
	}

	descriptionVO := NewOptionDescription(description)

	pollIDVO, err := NewPollID(pollID)
	if err != nil {
		return Option{}, err
	}

	optVotesVO, err := NewOptionVotes(votes)
	if err != nil {
		return Option{}, err
	}

	option := Option{
		id:          idVO,
		title:       titleVO,
		description: descriptionVO,
		pollID:      pollIDVO,
		votes:       optVotesVO,
	}

	return option, nil
}

// ID returns the Option identifier.
func (p Option) ID() OptionID {
	return p.id
}

// Title returns the Option title.
func (p Option) Title() OptionTitle {
	return p.title
}

// Description returns the Option description.
func (p Option) Description() OptionDescription {
	return p.description
}

// PollID returns the Option poll identifier.
func (p Option) PollID() PollID {
	return p.pollID
}

// Votes returns the Option votes.
func (p Option) Votes() OptionVotes {
	return p.votes
}

// DecreaseVotes decreases the Option votes.
func (p *Option) DecreaseVotes() error {
	if p.votes.value == 0 {
		return errors.NewWrongInput("cannot decrease votes when option has no votes")
	}

	p.votes.value--

	return nil
}

// IncreaseVotes increases the Option votes.
func (p *Option) IncreaseVotes() {
	p.votes.value++
}
