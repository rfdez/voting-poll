package creator_test

import (
	"context"
	"testing"

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

	creatorService := creator.NewService(pollRepositoryMock)

	err := creatorService.CreatePoll(context.Background(), pollID, pollTitle, pollDesc)

	pollRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_CreatePoll_Succeed(t *testing.T) {
	pollID, pollTitle, pollDesc := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description"

	pollRepositoryMock := new(storagemocks.PollRepository)
	pollRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("voting.Poll")).Return(nil)

	creatorService := creator.NewService(pollRepositoryMock)

	err := creatorService.CreatePoll(context.Background(), pollID, pollTitle, pollDesc)

	pollRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
