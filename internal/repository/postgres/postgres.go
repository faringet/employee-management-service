package postgres

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/zap"
)

type PostgresRepo struct {
	logger *zap.Logger
	JetDB  qrm.DB
}

func New(JetDB *sql.DB, logger *zap.Logger) (repository.Repository, error) {
	return &PostgresRepo{
		JetDB:  JetDB,
		logger: logger.Named("postgres"),
	}, nil
}

func DialectGoQu() goqu.DialectWrapper {
	return goqu.Dialect("postgres")
}
