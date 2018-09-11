package transport

import (
	"fmt"
	"context"
	"net/http"
)

func DecodeYouServiceFuncNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// deal request
	fmt.Println("request method  :", r.Method)
	return nil, nil
}

func EncodeYouServiceFuncNameResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// set response
	fmt.Println("Hello Visitor")
	return nil
}

func MakeYouServiceFuncNameEndpoint(ys YouServiceName) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// you can use you func  YouServiceFuncName
		ys.YouServiceFuncName()

		//return  you need
		return nil, nil
	}
}
