package increasing_test

import (
	"context"
	"testing"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/internal/increasing"
	"github.com/rfdez/voting-poll/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Service_IncreaseOptionVotes_RepositoryError(t *testing.T) {
	optID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	option, err := voting.NewOption(optID, "title", "description", "8aea44f4-50b9-421b-9eac-16ae6200ee32", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	optionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.OptionID")).Return(option, nil)
	optionRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Option")).Return(errors.New("error"))

	increasingService := increasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = increasingService.IncreaseOptionVotes(context.Background(), optID)

	optionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_IncreaseOptionVotes_Succeed(t *testing.T) {
	optID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	option, err := voting.NewOption(optID, "title", "description", "8aea44f4-50b9-421b-9eac-16ae6200ee32", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	optionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.OptionID")).Return(option, nil)
	optionRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Option")).Return(nil)

	increasingService := increasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = increasingService.IncreaseOptionVotes(context.Background(), optID)

	optionRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_Service_IncreasePollVoters_RepositoryError(t *testing.T) {
	pollID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	poll, err := voting.NewPoll(pollID, "title", "description", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(poll, nil)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(errors.New("error"))

	increasingService := increasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = increasingService.IncreasePollVoters(context.Background(), pollID)

	pollRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_IncreasePollVoters_Succeed(t *testing.T) {
	pollID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	poll, err := voting.NewPoll(pollID, "title", "description", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(poll, nil)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(nil)

	increasingService := increasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = increasingService.IncreasePollVoters(context.Background(), pollID)

	pollRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
