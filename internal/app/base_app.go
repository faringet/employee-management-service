package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/engagerocketco/go-common/auth0"
	"github.com/engagerocketco/go-common/ns"
	"github.com/nats-io/nats.go"

	"github.com/engagerocketco/go-common/config"
	"github.com/engagerocketco/go-common/pg"
	"github.com/engagerocketco/templates-api-svc/internal/handler"
	natsHandler "github.com/engagerocketco/templates-api-svc/internal/nats"
	"github.com/engagerocketco/templates-api-svc/internal/repository/postgres"
	natsClient "github.com/engagerocketco/templates-api-svc/internal/service/natsservice"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type BaseApp struct {
	cfg           *config.Config
	server        *handler.Server
	natsServer    *natsHandler.Server
	logger        *zap.Logger
	shutdownFuncs []func(context.Context) error
}

func NewBaseApp(ctx context.Context, logger *zap.Logger, cfg *config.Config) (*BaseApp, error) {
	b := &BaseApp{
		cfg:    cfg,
		logger: logger,
	}

	dbConn, err := pg.Connect(cfg.PostgresConfig.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to connect to the database: %w", err)
	}

	b.shutdownFuncs = append(b.shutdownFuncs, dbConn.Shutdown)

	natsConn, err := nats.Connect(cfg.NatsConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("nats: unable to connect to the nats server: %w", err)
	}

	b.natsServer = natsHandler.NewServer(natsConn, logger)
	b.shutdownFuncs = append(b.shutdownFuncs, b.natsServer.Shutdown)

	entityClient := ns.NewEntityClient(natsConn)
	permissionClient := ns.NewPermissionClient(natsConn)
	natsService := natsClient.NewNatsService(entityClient, permissionClient, logger)

	db, err := sql.Open("postgres", cfg.PostgresConfig.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to connect to the database: %w", err)
	}

	postgresRepo, err := postgres.New(db, logger)
	if err != nil {
		return nil, fmt.Errorf("repository: unable to initialize a postgres repository: %w", err)
	}

	templateService := templateservice.New(postgresRepo, natsService, logger)

	validateToken, err := auth0.EnsureValidToken(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("auth0: token validation: %w", err)
	}

	b.server = handler.NewServer(cfg, templateService, logger, validateToken)

	return b, nil
}

func (b *BaseApp) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := b.server.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			b.logger.Error("server run error", zap.Error(err))
			cancel()
		}
	}()

	b.logger.Info("server started", zap.String("addr", fmt.Sprintf("%s:%d", b.cfg.ServerHost, b.cfg.ServerPort)))

	select {
	case <-ctx.Done():
		b.logger.Error("stopping via context")
	case <-signalCh:
		b.logger.Info("service stopping")
	}

	cancel()

	for i := range b.shutdownFuncs {
		wg.Add(1)
		go func(f func(context.Context) error) {
			defer wg.Done()
			if err := f(context.Background()); err != nil {
				b.logger.Error("shutdown function failed", zap.Error(err))
			}
		}(b.shutdownFuncs[i])
	}

	wg.Wait()

	b.logger.Info("service stopped")
}
