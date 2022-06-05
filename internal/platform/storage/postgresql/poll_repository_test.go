package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/errors"
	"github.com/rfdez/voting-poll/internal/platform/storage/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PollRepository_Save_RepositoryError(t *testing.T) {
	pollID, pollTitle, pollDesc := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description"
	poll, err := voting.NewPoll(pollID, pollTitle, pollDesc)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO polls (id, title, description) VALUES ($1, $2, $3)").
		WithArgs(pollID, pollTitle, pollDesc).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewPollRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), poll)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_PollRepository_Save_Succeed(t *testing.T) {
	pollID, pollTitle, pollDesc := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description"
	poll, err := voting.NewPoll(pollID, pollTitle, pollDesc)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO polls (id, title, description) VALUES ($1, $2, $3)").
		WithArgs(pollID, pollTitle, pollDesc).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := postgresql.NewPollRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), poll)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
