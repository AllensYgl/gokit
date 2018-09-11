package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"log"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"

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

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc), //frist start makeUppercaseEndpoint  	//third start func
		decodeUppercaseRequest,     // second start
		encodeResponse,             // last start
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	//fmt.Println(uppercaseHandler, countHandler)
	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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
