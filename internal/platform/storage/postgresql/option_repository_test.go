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
	id, title, desc, pollID := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591"
	option, err := voting.NewOption(id, title, desc, pollID)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO options (id, title, description, poll_id) VALUES ($1, $2, $3, $4)").
		WithArgs(id, title, desc, pollID).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewOptionRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), option)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_OptionRepository_Save_Succeed(t *testing.T) {
	id, title, desc, pollID := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Poll", "Test description", "fbe97bf6-0a94-4ca7-90cc-b27361318591"
	option, err := voting.NewOption(id, title, desc, pollID)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO options (id, title, description, poll_id) VALUES ($1, $2, $3, $4)").
		WithArgs(id, title, desc, pollID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := postgresql.NewOptionRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), option)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
