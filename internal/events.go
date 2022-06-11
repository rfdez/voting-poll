package voting

import (
	"time"

	"github.com/rfdez/voting-poll/kit/event"
)

const (
	VoteDeletedEventType event.Type = "voting-app.voting-vote.1.event.vote.deleted"
)

type VoteDeletedEvent struct {
	event.BaseEvent
	pollID   string
	optionID string
}

func NewVoteDeletedEvent(pollID, optionID string) VoteDeletedEvent {
	return VoteDeletedEvent{
		pollID:   pollID,
		optionID: optionID,

		BaseEvent: event.NewBaseEvent((pollID + "-" + optionID), "", time.Now()),
	}
}

func (e VoteDeletedEvent) Type() event.Type {
	return VoteDeletedEventType
}

func (e VoteDeletedEvent) FromPrimitives(aggregateID string, body map[string]interface{}, id, occurredOn string) (event.Event, error) {
	eventOccurredOn, err := time.Parse(time.RFC3339, occurredOn)
	if err != nil {
		return nil, err
	}

	return VoteDeletedEvent{
		pollID:    body["poll_id"].(string),
		optionID:  body["option_id"].(string),
		BaseEvent: event.NewBaseEvent(id, aggregateID, eventOccurredOn),
	}, nil
}

func (e VoteDeletedEvent) ToPrimitives() map[string]interface{} {
	return map[string]interface{}{
		"poll_id":   e.pollID,
		"option_id": e.optionID,
	}
}

func (e VoteDeletedEvent) PollID() string {
	return e.pollID
}

func (e VoteDeletedEvent) OptionID() string {
	return e.optionID
}
