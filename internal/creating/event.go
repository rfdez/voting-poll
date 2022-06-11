package creating

import (
	"context"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/increasing"
	"github.com/rfdez/voting-poll/kit/event"
)

type IncreaseOptionVotesOnVoteCreated struct {
	increasingService increasing.Service
}

func NewIncreaseOptionVotesOnVoteCreated(increasingService increasing.Service) IncreaseOptionVotesOnVoteCreated {
	return IncreaseOptionVotesOnVoteCreated{
		increasingService: increasingService,
	}
}

func (e IncreaseOptionVotesOnVoteCreated) Handle(ctx context.Context, evt event.Event) error {
	voteDeleteEvt, ok := evt.(voting.VoteCreatedEvent)
	if !ok {
		return nil
	}

	return e.increasingService.IncreaseOptionVotes(ctx, voteDeleteEvt.OptionID())
}

type IncreasePollVotersOnVoteCreated struct {
	increasingService increasing.Service
}

func NewIncreasePollVotersOnVoteCreated(increasingService increasing.Service) IncreasePollVotersOnVoteCreated {
	return IncreasePollVotersOnVoteCreated{
		increasingService: increasingService,
	}
}

func (e IncreasePollVotersOnVoteCreated) Handle(ctx context.Context, evt event.Event) error {
	voteDeleteEvt, ok := evt.(voting.VoteCreatedEvent)
	if !ok {
		return nil
	}

	return e.increasingService.IncreasePollVoters(ctx, voteDeleteEvt.PollID())
}
