package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/engagerocketco/go-common/config"
	testsUtils "github.com/engagerocketco/go-common/tests"
	"github.com/engagerocketco/templates-api-svc/internal/repository"
	"github.com/engagerocketco/templates-api-svc/internal/repository/postgres"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func NewPgRepo(
	ctx context.Context,
	t *testing.T,
	cfg *config.Config,
) (repository.Repository, *testsUtils.Conn) {
	conn := testsUtils.NewTestPgConn(ctx, t, cfg)
	db, err := sqlx.Connect("postgres", cfg.PostgresConfig.ConnectionString())
	if err != nil {
		t.Fatalf("create pg conn: %v", err)
	}

	jetDB, err := sql.Open("postgres", cfg.PostgresConfig.ConnectionString())
	if err != nil {
		t.Fatalf("create pg conn: %v", err)
	}

	repo, err := postgres.New(db, jetDB, &zap.Logger{})
	if err != nil {
		t.Fatalf("new postgres: %v", err)
	}

	return repo, conn
}
