package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/errors"
)

type pollRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewPollRepository instantiate the PollRepository
func NewPollRepository(db *sql.DB, dbTimeout time.Duration) voting.PollRepository {
	return &pollRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Find implements the PollRepository interface.
func (r *pollRepository) Find(ctx context.Context, id voting.PollID) (voting.Poll, error) {
	pollSQLStruct := sqlbuilder.NewStruct(new(sqlPoll))

	sb := pollSQLStruct.SelectFrom(sqlPollTable)
	sb.Where(sb.E("id", id.String()))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var poll sqlPoll
	if err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(pollSQLStruct.Addr(&poll)...); err != nil {
		if err == sql.ErrNoRows {
			return voting.Poll{}, errors.NewNotFound("poll %s not found", id.String())
		}

		return voting.Poll{}, errors.Wrap(err, "error finding poll")
	}

	p, err := voting.NewPoll(poll.ID, poll.Title, poll.Description)
	if err != nil {
		return voting.Poll{}, err
	}

	return p, nil
}

// Save implements the PollRepository interface.
func (r *pollRepository) Save(ctx context.Context, poll voting.Poll) error {
	pollSQLStruct := sqlbuilder.NewStruct(new(sqlPoll))

	sb := pollSQLStruct.InsertInto(sqlPollTable, sqlPoll{
		ID:          poll.ID().String(),
		Title:       poll.Title().String(),
		Description: poll.Description().String(),
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return errors.Wrap(err, "error saving poll")
	}

	return nil
}
