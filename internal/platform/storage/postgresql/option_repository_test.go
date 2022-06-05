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

func Test_OptionRepository_Save_RepositoryError(t *testing.T) {
	id, title, desc, pollID, votes := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591", 0
	option, err := voting.NewOption(id, title, desc, pollID, votes)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(`
		INSERT INTO options (id, title, description, poll_id, votes) VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (id)
			DO UPDATE SET
				title = EXCLUDED.title, description = EXCLUDED.description, poll_id = EXCLUDED.poll_id, votes = EXCLUDED.votes`).
		WithArgs(id, title, desc, pollID, votes).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewOptionRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), option)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_OptionRepository_Save_Succeed(t *testing.T) {
	id, title, desc, pollID, votes := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591", 0
	option, err := voting.NewOption(id, title, desc, pollID, votes)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(`
		INSERT INTO options (id, title, description, poll_id, votes) VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (id)
			DO UPDATE SET
				title = EXCLUDED.title, description = EXCLUDED.description, poll_id = EXCLUDED.poll_id, votes = EXCLUDED.votes`).
		WithArgs(id, title, desc, pollID, votes).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := postgresql.NewOptionRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), option)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func Test_OptionRepository_Find_RepositoryError(t *testing.T) {
	id, err := voting.NewOptionID("37a0f027-15e6-47cc-a5d2-64183281087e")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT options.id, options.title, options.description, options.poll_id, options.votes FROM options WHERE id = $1").
		WithArgs(id.String()).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewOptionRepository(db, 1*time.Millisecond)

	_, err = repo.Find(context.Background(), id)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_OptionRepository_Find_Succeed(t *testing.T) {
	id, err := voting.NewOptionID("37a0f027-15e6-47cc-a5d2-64183281087e")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "title", "description", "poll_id", "votes"}).
		AddRow(id.String(), "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591", 0)

	sqlMock.ExpectQuery(
		"SELECT options.id, options.title, options.description, options.poll_id, options.votes FROM options WHERE id = $1").
		WithArgs(id.String()).
		WillReturnRows(rows)

	repo := postgresql.NewOptionRepository(db, 1*time.Millisecond)

	opt, err := repo.Find(context.Background(), id)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, id.String(), opt.ID().String())
}
