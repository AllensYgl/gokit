package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

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
