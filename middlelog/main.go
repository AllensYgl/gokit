package main

import (
	"net/http"
	"os"
	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)


func main() {
	logger := log.NewLogfmtLogger(os.Stderr)


	//inherit		继承
	//polymorphism	多态
	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	uppercaseHandlers := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),	//	func (mw loggingMiddleware) Uppercase
		decodeUppercaseRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/", uppercaseHandlers)
	http.ListenAndServe(":8080", nil)
}
