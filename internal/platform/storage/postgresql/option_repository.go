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

// Save implements the OptionRepository interface.
func (r *optionRepository) Save(ctx context.Context, option voting.Option) error {
	optSQLStruct := sqlbuilder.NewStruct(new(sqlOption))

	sb := optSQLStruct.InsertInto(sqlOptionTable, sqlOption{
		ID:          option.ID().String(),
		Title:       option.Title().String(),
		Description: option.Description().String(),
		PollID:      option.PollID().String(),
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return errors.Wrap(err, "error saving option")
	}

	return nil
}
