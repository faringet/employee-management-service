package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type PingResponse string

// Profile godoc
// Test
//
//	@Summary		Ping Request
//	@Description	Ping Request
//	@Tags			ping
//	@Produce		plain
//
//	@Success		200	{string}	string	"pong"
//
//	@Router			/ping [get]
func MakePingEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "pong", nil
	}
}
