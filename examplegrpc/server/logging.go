package main

import (
	"golang.org/x/net/context"
	"gokit/examplegrpc/pb"
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next *grpc.Server) *grpc.Server {
		return &logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	pb.YouServiceNameServer
}

func (mw *logmw) YouServiceFuncName(ctx context.Context, in *pb.ServiceRequest) (rep *pb.ServiceReply, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", in.Name,
			"output", rep.Message,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	rep, err = mw.YouServiceNameServer.YouServiceFuncName(ctx,in)
	return
}
