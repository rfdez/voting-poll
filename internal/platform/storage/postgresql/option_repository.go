package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	voting "github.com/rfdez/voting-poll/internal"
	"github.com/rfdez/voting-poll/internal/errors"
)

type optionRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewOptionRepository instantiate the OptionRepository
func NewOptionRepository(db *sql.DB, dbTimeout time.Duration) voting.OptionRepository {
	return &optionRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Find implements the OptionRepository interface.
func (r *optionRepository) Find(ctx context.Context, id voting.OptionID) (voting.Option, error) {
	optSQLStruct := sqlbuilder.NewStruct(new(sqlOption))

	sb := optSQLStruct.SelectFrom(sqlOptionTable)
	sb.Where(sb.E("id", id.String()))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var opt sqlOption
	if err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(optSQLStruct.Addr(&opt)...); err != nil {
		if err == sql.ErrNoRows {
			return voting.Option{}, errors.NewNotFound("option %s not found", id.String())
		}

		return voting.Option{}, errors.Wrap(err, "error finding option")
	}

	p, err := voting.NewOption(opt.ID, opt.Title, opt.Description, opt.PollID, opt.Votes)
	if err != nil {
		return voting.Option{}, err
	}

	return p, nil
}

// Save implements the OptionRepository interface.
func (r *optionRepository) Save(ctx context.Context, option voting.Option) error {
	optSQLStruct := sqlbuilder.NewStruct(new(sqlOption))

	sb := optSQLStruct.InsertInto(sqlOptionTable, sqlOption{
		ID:          option.ID().String(),
		Title:       option.Title().String(),
		Description: option.Description().String(),
		PollID:      option.PollID().String(),
		Votes:       option.Votes().Value(),
	})
	sb.SQL(`
		ON CONFLICT (id)
		DO UPDATE SET
			title = EXCLUDED.title, description = EXCLUDED.description, poll_id = EXCLUDED.poll_id, votes = EXCLUDED.votes`)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return errors.Wrap(err, "error saving option")
	}

	return nil
}
