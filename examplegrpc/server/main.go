package main

import (
	"golang.org/x/net/context"
	"gokit/examplegrpc/pb"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) YouServiceFuncName(ctx context.Context, in *pb.ServiceRequest) (*pb.ServiceReply, error) {
	return &pb.ServiceReply{Message: "Hello " + in.Name}, nil
}

type ServiceMiddleware func(pb.YouServiceNameServer) pb.YouServiceNameServer

func makeYouServiceFuncNameEndpoint(ys server) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// you can use you func  YouServiceFuncName
		req:=request.(*pb.ServiceRequest)
		ys.YouServiceFuncName(ctx,req)
		//return  you need
		return nil, nil
	}
}

func decodeYouServiceFuncNameRequest(_ context.Context, req *pb.ServiceRequest) (interface{}, error) {
	// deal request
	//fmt.Println("request method  :", r.Method)
	return nil, nil
}

func encodeYouServiceFuncNameResponse(_ context.Context, rep *pb.ServiceReply, response interface{}) error {
	// set response
	//fmt.Println("Hello Visitor")
	return nil
}

func main() {
	ys := server{}
	//	2.	bound func
	//	decodeYouServiceFuncNameRequest		run   get a endpoint.Endpoint
	printVisitorNameHandler := httptransport.NewServer(
		makeYouServiceFuncNameEndpoint(ys),
		decodeYouServiceFuncNameRequest,
		encodeYouServiceFuncNameResponse,
	)

	// 	3.	set router
	http.Handle("/", printVisitorNameHandler)

	// 	4.	set listen port
	log.Fatal(http.ListenAndServe(":8080", nil))
}
