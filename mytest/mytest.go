package main

import (
	"context"
	"fmt"
	"log"

	"net/http"

	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

type MyService interface {
	PrintVisitorName()
}

type myService struct {
	serviceName string
	option      string
}

func (ms myService) PrintVisitorName() {
	fmt.Println("serviceName:", ms.serviceName)
	fmt.Println("optionName:", ms.option)
}

func makePrintVisitorNameEndpoint(ms MyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ms.PrintVisitorName()
		return nil, nil
	}
}

func main() {
	ms := myService{serviceName: "myservice", option: "print"}

	printVisitorNameHandler := httptransport.NewServer(
		makePrintVisitorNameEndpoint(ms),
		decodePrintVisitorNameRequest,
		encodeResponse,
	)

	http.Handle("/", printVisitorNameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodePrintVisitorNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("request method:", r.Method)
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("Hello Visitor")
	return nil
}
