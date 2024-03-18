package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Conn struct {
	*sqlx.DB
}

func Connect(connectionString string) (*Conn, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	dbx := sqlx.NewDb(db, "postgres")
	return &Conn{
		dbx,
	}, nil
}

func (conn *Conn) Shutdown(ctx context.Context) error {
	return conn.Close()
}
