package creator_test

import (
	"context"
	"testing"

	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/creator"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Service_CreatePoll_RepositoryError(t *testing.T) {
	pollID, pollTitle, pollDesc := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description"

	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(errors.New("error"))
	optionRepositoryMock := new(storagemocks.OptionRepository)

	creatorService := creator.NewService(pollRepositoryMock, optionRepositoryMock)

	err := creatorService.CreatePoll(context.Background(), pollID, pollTitle, pollDesc)

	pollRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_CreatePoll_Succeed(t *testing.T) {
	pollID, pollTitle, pollDesc := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description"

	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(nil)
	optionRepositoryMock := new(storagemocks.OptionRepository)

	creatorService := creator.NewService(pollRepositoryMock, optionRepositoryMock)

	err := creatorService.CreatePoll(context.Background(), pollID, pollTitle, pollDesc)

	pollRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_Service_CreateOption_RepositoryError(t *testing.T) {
	id, optionTitle, optionDesc, pollID := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591"

	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(voting.Poll{}, nil)
	optionRepositoryMock := new(storagemocks.OptionRepository)
	optionRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Option")).Return(errors.New("error"))

	creatorService := creator.NewService(pollRepositoryMock, optionRepositoryMock)

	err := creatorService.CreateOption(context.Background(), id, optionTitle, optionDesc, pollID)

	pollRepositoryMock.AssertExpectations(t)
	optionRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_CreateOption_Succeed(t *testing.T) {
	id, optionTitle, optionDesc, pollID := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591"

	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Find", mock.Anything, mock.AnythingOfType("voting.PollID")).Return(voting.Poll{}, nil)
	optionRepositoryMock := new(storagemocks.OptionRepository)
	optionRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Option")).Return(nil)

	creatorService := creator.NewService(pollRepositoryMock, optionRepositoryMock)

	err := creatorService.CreateOption(context.Background(), id, optionTitle, optionDesc, pollID)

	pollRepositoryMock.AssertExpectations(t)
	optionRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
