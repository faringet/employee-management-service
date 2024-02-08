package postgres

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	"go.uber.org/zap"
)

type PostgresRepo struct {
	DB     goqu.SQLDatabase
	logger *zap.Logger
	JetDB  *sql.DB
}

func New(db goqu.SQLDatabase, JetDB *sql.DB, logger *zap.Logger) (repository.Repository, error) {
	return &PostgresRepo{
		DB:     db,
		JetDB:  JetDB,
		logger: logger.Named("postgres"),
	}, nil
}

func DialectGoQu() goqu.DialectWrapper {
	return goqu.Dialect("postgres")
}
