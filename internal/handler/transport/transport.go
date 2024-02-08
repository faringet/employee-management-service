package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	ie "github.com/engagerocketco/templates-api-svc/pkg/errors"
)

var transportLogger *zap.Logger
var finalyzerLogger *zap.Logger

func NewTransportLogger(logger *zap.Logger) {
	transportLogger = logger.Named("handler").Named("transport")
	finalyzerLogger = logger.Named("handler").Named("finalyzer")
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	var responseError ie.Error
	errors.As(err, &responseError)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(responseError.Code)
	err = json.NewEncoder(w).Encode(responseError)
	if err != nil {
		transportLogger.Error("error encode responseError", zap.Error(err))
	}
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func finalyzer(ctx context.Context, code int, r *http.Request) {
	finalyzerLogger.Info("request",
		zap.String("path", r.URL.Path),
		zap.String("method", r.Method),
		zap.Int("code", code),
	)
}
