package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

// option interface about string
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

// option struct
type stringService struct{}

// func realize
func (str stringService) Uppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (str stringService) Count(ctx context.Context, s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")

// define struct about request and response
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// point  define
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// adapter
// get a option function Endpoint
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	fmt.Println("bbb")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("ccc")
		req := request.(uppercaseRequest)   //	get a option struct
		v, err := svc.Uppercase(ctx, req.S) // use option at service in this func
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(ctx, req.S)
		return countResponse{v}, nil
	}

}

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

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

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("aaa")
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("ddd")
	return json.NewEncoder(w).Encode(response)
}
