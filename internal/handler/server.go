package handler

import (
	"context"
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/engagerocketco/go-common/config"
	"github.com/engagerocketco/templates-api-svc/docs"
	"github.com/engagerocketco/templates-api-svc/internal/handler/transport"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	*http.Server
}

func NewServer(
	cfg *config.Config,
	templateService templateservice.Service,
	logger *zap.Logger,
	middlewares ...func(next http.Handler) http.Handler,
) *Server {
	srv := Server{
		&http.Server{Addr: fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)},
	}

	docs.SwaggerInfo.Host = cfg.SwaggerConfig.Host

	transport.NewTransportLogger(logger)

	r := mux.NewRouter().PathPrefix("/api/v1/template").Subrouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	srv.Handler = r
	return &srv
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			errCh <- fmt.Errorf("server: failed to listen and serve: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		return s.Shutdown(context.Background())
	case err := <-errCh:
		return err
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
