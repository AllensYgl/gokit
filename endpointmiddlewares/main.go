package main

import (
	"context"
	"net/http"
	"os"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

// to show Transport logging
func loggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}

func loggingMiddle(logger log.Logger, next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.Log("msg", "calling endpoint")
		defer logger.Log("msg", "called endpoint")
		return next(ctx, request)
	}
}

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	svc := stringService{}

	// point
	var uppercase endpoint.Endpoint
	uppercase = makeUppercaseEndpoint(svc)
	uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

	var upper endpoint.Endpoint
	upper = loggingMiddle(log.With(logger, "method", "uppercase"), makeUppercaseEndpoint(svc))

	uppercaseHandlers := httptransport.NewServer(
		upper,
		decodeUppercaseRequest,
		encodeResponse,
	)
	uppercaseHandler := httptransport.NewServer(
		uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/", uppercaseHandlers)
	http.ListenAndServe(":8080", nil)
}