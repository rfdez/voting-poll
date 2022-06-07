package uuid

import (
	"github.com/google/uuid"
	"github.com/rfdez/voting-poll/internal/errors"
)

func New(s string) (string, error) {
	v, err := uuid.Parse(s)
	if err != nil {
		return "", errors.WrapWrongInput(err, "invalid uuid %s", s)
	}

	return v.String(), nil
}

func Generate() string {
	return uuid.New().String()
}
