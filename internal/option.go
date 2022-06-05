package voting

import (
	"context"

	"github.com/google/uuid"
	"github.com/rfdez/voting-poll/internal/errors"
)

// OptionID represents a option identifier.
type OptionID struct {
	value string
}

// NewOptionID instantiate the VO for OptionID.
func NewOptionID(value string) (OptionID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return OptionID{}, errors.WrapWrongInput(err, "invalid option id %s", value)
	}

	return OptionID{
		value: v.String(),
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

// Option is the data structure that represents a poll option.
type Option struct {
	id          OptionID
	title       OptionTitle
	description OptionDescription
	pollID      PollID
}

// OptionRepository is the interface that must be implemented by the Option repository.
type OptionRepository interface {
	Save(context.Context, Option) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=OptionRepository

// NewOption creates a new Option.
func NewOption(id, name, description, pollID string) (Option, error) {
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

	option := Option{
		id:          idVO,
		title:       titleVO,
		description: descriptionVO,
		pollID:      pollIDVO,
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
