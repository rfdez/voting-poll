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
	id, title, desc, voters := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", 0
	poll, err := voting.NewPoll(id, title, desc, voters)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(`
		INSERT INTO polls (id, title, description, voters) VALUES ($1, $2, $3, $4)
			ON CONFLICT (id)
			DO UPDATE SET
				title = EXCLUDED.title, description = EXCLUDED.description, voters = EXCLUDED.voters`).
		WithArgs(id, title, desc, voters).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewPollRepository(db, 5*time.Second)

	err = repo.Save(context.Background(), poll)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_PollRepository_Save_Succeed(t *testing.T) {
	id, title, desc, voters := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", 0
	poll, err := voting.NewPoll(id, title, desc, voters)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(`
		INSERT INTO polls (id, title, description, voters) VALUES ($1, $2, $3, $4)
			ON CONFLICT (id)
			DO UPDATE SET
				title = EXCLUDED.title, description = EXCLUDED.description, voters = EXCLUDED.voters`).
		WithArgs(id, title, desc, voters).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := postgresql.NewPollRepository(db, 5*time.Second)

	err = repo.Save(context.Background(), poll)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func Test_PollRepository_Find_RepositoryError(t *testing.T) {
	id, err := voting.NewPollID("37a0f027-15e6-47cc-a5d2-64183281087e")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT polls.id, polls.title, polls.description, polls.voters FROM polls WHERE id = $1").
		WithArgs(id.String()).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewPollRepository(db, 5*time.Second)

	_, err = repo.Find(context.Background(), id)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_PollRepository_Find_Succeed(t *testing.T) {
	id, err := voting.NewPollID("37a0f027-15e6-47cc-a5d2-64183281087e")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "voters"}).
		AddRow(id.String(), "Test Poll", "Test description", 0)

	sqlMock.ExpectQuery(
		"SELECT polls.id, polls.title, polls.description, polls.voters FROM polls WHERE id = $1").
		WithArgs(id.String()).
		WillReturnRows(rows)

	repo := postgresql.NewPollRepository(db, 5*time.Second)

	poll, err := repo.Find(context.Background(), id)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, id.String(), poll.ID().String())
}
