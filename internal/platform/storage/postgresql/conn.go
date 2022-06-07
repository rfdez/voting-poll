package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConnection(user, pass, host string, port uint, dbName, params string) (*sql.DB, error) {
	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", user, pass, host, port, dbName, params)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}
