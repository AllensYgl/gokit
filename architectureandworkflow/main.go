package main

import (
	"net/http"
	"gokit/architecture/service"
	"gokit/architecture/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
)


func main() {

	//workflow
	// 	1.	define service struct
	ys := service.YouServiceStruct{}

	//	2.	bound func
	//	decodeYouServiceFuncNameRequest		run   get a endpoint.Endpoint
	printVisitorNameHandler := httptransport.NewServer(
		transport.MakeYouServiceFuncNameEndpoint(ys),
		transport.DecodeYouServiceFuncNameRequest,
		transport.EncodeYouServiceFuncNameResponse,
	)

	// 	3.	set router
	http.Handle("/", printVisitorNameHandler)

	// 	4.	set listen port
	log.Fatal(http.ListenAndServe(":8080", nil))

	// 	5.	request ------->router
	// 	6.	endpoint.Endpoint   				run
	//	7.	makeYouServiceFuncNameEndpoint    	run
	// 	8.	encodeYouServiceFuncNameResponse	run

}
