package decreasing_test

import (
	"context"
	"testing"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/decreasing"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Service_DecreaseOptionVotes_RepositoryError(t *testing.T) {
	optID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	option, err := voting.NewOption(optID, "title", "description", "8aea44f4-50b9-421b-9eac-16ae6200ee32", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	optionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.OptionID")).Return(option, nil)
	optionRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Option")).Return(errors.New("error"))

	decreasingService := decreasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = decreasingService.DecreaseOptionVotes(context.Background(), optID)

	optionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_DecreaseOptionVotes_DecreaseError(t *testing.T) {
	optID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	optVotes := 0
	option, err := voting.NewOption(optID, "title", "description", "8aea44f4-50b9-421b-9eac-16ae6200ee32", optVotes)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	optionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.OptionID")).Return(option, nil)

	decreasingService := decreasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = decreasingService.DecreaseOptionVotes(context.Background(), optID)

	optionRepositoryMock.AssertNotCalled(t, "Save", mock.Anything, mock.AnythingOfType("voting.Option"))
	optionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_DecreaseOptionVotes_Succeed(t *testing.T) {
	optID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	option, err := voting.NewOption(optID, "title", "description", "8aea44f4-50b9-421b-9eac-16ae6200ee32", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	optionRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.OptionID")).Return(option, nil)
	optionRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Option")).Return(nil)

	decreasingService := decreasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = decreasingService.DecreaseOptionVotes(context.Background(), optID)

	optionRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_Service_DecreasePollVoters_RepositoryError(t *testing.T) {
	pollID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	poll, err := voting.NewPoll(pollID, "title", "description", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(poll, nil)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(errors.New("error"))

	decreasingService := decreasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = decreasingService.DecreasePollVoters(context.Background(), pollID)

	pollRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_DecreasePollVoters_DecreaseError(t *testing.T) {
	pollID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	pollVoters := 0
	poll, err := voting.NewPoll(pollID, "title", "description", pollVoters)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(poll, nil)

	decreasingService := decreasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = decreasingService.DecreasePollVoters(context.Background(), pollID)

	pollRepositoryMock.AssertNotCalled(t, "Save", mock.Anything, mock.AnythingOfType("voting.Poll"))
	pollRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_DecreasePollVoters_Succeed(t *testing.T) {
	pollID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	poll, err := voting.NewPoll(pollID, "title", "description", 2)
	require.NoError(t, err)

	optionRepositoryMock := new(storagemocks.OptionRepository)
	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(poll, nil)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(nil)

	decreasingService := decreasing.NewService(pollRepositoryMock, optionRepositoryMock)

	err = decreasingService.DecreasePollVoters(context.Background(), pollID)

	pollRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
