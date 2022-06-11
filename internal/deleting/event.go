package deleting

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/decreasing"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/kit/event"
)

type DecreaseOptionVotesOnVoteDeleted struct {
	decreasingService decreasing.Service
}

func NewDecreaseOptionVotesOnVoteDeleted(decreasingService decreasing.Service) DecreaseOptionVotesOnVoteDeleted {
	return DecreaseOptionVotesOnVoteDeleted{
		decreasingService: decreasingService,
	}
}

func (e DecreaseOptionVotesOnVoteDeleted) Handle(ctx context.Context, evt event.Event) error {
	voteDeleteEvt, ok := evt.(voting.VoteDeletedEvent)
	if !ok {
		return nil
	}

	return e.decreasingService.DecreaseOptionVotes(ctx, voteDeleteEvt.OptionID())
}

type DecreasePollVotersOnVoteDeleted struct {
	decreasingService decreasing.Service
}

func NewDecreasePollVotersOnVoteDeleted(decreasingService decreasing.Service) DecreasePollVotersOnVoteDeleted {
	return DecreasePollVotersOnVoteDeleted{
		decreasingService: decreasingService,
	}
}

func (e DecreasePollVotersOnVoteDeleted) Handle(ctx context.Context, evt event.Event) error {
	voteDeleteEvt, ok := evt.(voting.VoteDeletedEvent)
	if !ok {
		return errors.NewWrongInput("unknown event type")
	}

	return e.decreasingService.DecreasePollVoters(ctx, voteDeleteEvt.PollID())
}
