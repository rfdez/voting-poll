package voting

import "github.com/rfdez/voting-poll/kit/event"

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

		BaseEvent: event.NewBaseEvent(pollID + "-" + optionID),
	}
}

func (e VoteDeletedEvent) Type() event.Type {
	return VoteDeletedEventType
}

func (e VoteDeletedEvent) PollID() string {
	return e.pollID
}

func (e VoteDeletedEvent) OptionID() string {
	return e.optionID
}
