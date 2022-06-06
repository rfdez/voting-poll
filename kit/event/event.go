package event

import (
	"context"
	"time"

	"github.com/rfdez/voting-poll/kit/uuid"
)

// Bus defines the expected behavior of an event bus.
type Bus interface {
	// Publish publishes an event to the bus.
	Publish(context.Context, Event) error
	// Subscribe subscribes to events of the given type.
	Subscribe(Type, Handler) error
}

//go:generate mockery --case=snake --outpkg=eventmocks --output=eventmocks --name=Bus

// Handler defines the expected behavior of an event handler.
type Handler interface {
	// Handle handles an event.
	Handle(context.Context, Event) error
}

// Type represents the type of an event.
type Type string

// Event represents an event.
type Event interface {
	// ID returns the ID of the event.
	ID() string
	// AggregateID returns the ID of the aggregate the event belongs to.
	AggregateID() string
	// OccurredOn returns the time the event occurred on.
	OccurredOn() time.Time
	// Type returns the type of the event.
	Type() Type
}

type BaseEvent struct {
	id          string
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		id:          uuid.Generate(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (e BaseEvent) ID() string {
	return e.id
}

func (e BaseEvent) AggregateID() string {
	return e.aggregateID
}

func (e BaseEvent) OccurredOn() time.Time {
	return e.occurredOn
}
