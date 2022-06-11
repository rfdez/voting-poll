package voting

import (
	"context"

	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/kit/uuid"
)

// PollID represents a poll identifier.
type PollID struct {
	value string
}

// NewPollID instantiate the VO for PollID.
func NewPollID(value string) (PollID, error) {
	v, err := uuid.New(value)
	if err != nil {
		return PollID{}, err
	}

	return PollID{
		value: v,
	}, nil
}

// String returns the string representation of the PollID.
func (id PollID) String() string {
	return id.value
}

// PollTitle represents a poll title.
type PollTitle struct {
	value string
}

// NewPollTitle instantiate the VO for PollTitle.
func NewPollTitle(value string) (PollTitle, error) {
	if value == "" {
		return PollTitle{}, errors.NewWrongInput("poll title cannot be empty")
	}

	return PollTitle{
		value: value,
	}, nil
}

// String returns the string representation of the PollTitle.
func (title PollTitle) String() string {
	return title.value
}

// PollDescription represents a poll description.
type PollDescription struct {
	value string
}

// NewPollDescription instantiate the VO for PollDescription.
func NewPollDescription(value string) PollDescription {
	if value == "" {
		value = "No description provided"
	}

	return PollDescription{
		value: value,
	}
}

// String returns the string representation of the PollDescription.
func (description PollDescription) String() string {
	return description.value
}

// PollVoters represents the number of voters for a poll.
type PollVoters struct {
	value int
}

// NewPollVoters creates a new PollVoters.
func NewPollVoters(value int) (PollVoters, error) {
	if value < 0 {
		return PollVoters{}, errors.NewWrongInput("poll voters cannot be negative")
	}

	return PollVoters{
		value: value,
	}, nil
}

// Value returns the Option votes.
func (votes PollVoters) Value() int {
	return votes.value
}

// Poll is the data structure that represents a poll.
type Poll struct {
	id          PollID
	title       PollTitle
	description PollDescription
	voters      PollVoters
}

// PollRepository is the interface that must be implemented by the poll repository.
type PollRepository interface {
	Find(context.Context, PollID) (Poll, error)
	Save(context.Context, Poll) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=PollRepository

// NewPoll creates a new poll.
func NewPoll(id, name, description string, voters int) (Poll, error) {
	idVO, err := NewPollID(id)
	if err != nil {
		return Poll{}, err
	}

	titleVO, err := NewPollTitle(name)
	if err != nil {
		return Poll{}, err
	}

	descriptionVO := NewPollDescription(description)

	votersVO, err := NewPollVoters(voters)
	if err != nil {
		return Poll{}, err
	}

	poll := Poll{
		id:          idVO,
		title:       titleVO,
		description: descriptionVO,
		voters:      votersVO,
	}

	return poll, nil
}

// ID returns the poll identifier.
func (p Poll) ID() PollID {
	return p.id
}

// Title returns the poll title.
func (p Poll) Title() PollTitle {
	return p.title
}

// Description returns the poll description.
func (p Poll) Description() PollDescription {
	return p.description
}

// Voters returns the Poll voters.
func (p Poll) Voters() PollVoters {
	return p.voters
}

// DecreaseVoters decreases the Poll voters.
func (p *Poll) DecreaseVoters() error {
	if p.voters.value == 0 {
		return errors.NewWrongInput("cannot decrease votes when poll has no voters")
	}

	p.voters.value--

	return nil
}
