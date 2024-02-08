package nats

import (
	"context"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Server struct {
	conn *nats.Conn
}

func NewServer(conn *nats.Conn, logger *zap.Logger) *Server {

	return &Server{
		conn: conn,
	}
}

func (n *Server) Shutdown(ctx context.Context) error {
	return n.conn.Drain()
}
