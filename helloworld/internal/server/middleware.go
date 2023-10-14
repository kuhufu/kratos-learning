package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"helloworld/api/errs"
	v1 "helloworld/api/helloworld/v1"
	http2 "net/http"
	"time"
)

func httpLog(logger log.Logger) middleware.Middleware {
	helper := log.NewHelper(logger, log.WithMessageKey("code"))

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				method string
				path   string
				arg    = req
				code   = int32(200)
				reason string
				uid    int64
				//kind      string
				//operation string
				//level log.Level = log.LevelInfo
				logMsg string
				msg    string
			)
			startTime := time.Now()
			time.Sleep(time.Millisecond)
			request, _ := http.RequestFromServerContext(ctx)
			method = request.Method
			path = request.RequestURI
			reply, err = handler(ctx, req)

			kvs := []any{
				"method", request.Method,
				"path", request.RequestURI,
				"args", req,
			}

			if se := errors.FromError(err); se != nil {
				code = errs.ErrorReason_value[se.Reason]
				reason = se.Reason
				msg = se.Message
				//level = log.LevelError
				kvs = append(kvs, []any{
					"reason", reason,
					"message", se.Message,
				}...)
			}

			kvs = append(kvs, []any{
				"latency", time.Since(startTime),
			}...)

			if err != nil {
				helper.Errorf("  %-6v  uid=%-6v %v %v %v %v:%v %v", code, uid, method, path, arg, reason, msg, logMsg)
			} else {
				helper.Infof("  %-6v  uid=%-6v %v %v %v %v:%v", code, uid, method, path, arg)
			}
			return
		}
	}
}

func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}

func relyErrorMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		if err == nil {
			return
		}

		fromError := errors.FromError(err)
		if fromError != nil {
			fromError.Code = errs.ErrorReason_value[fromError.Reason]
			return nil, fromError
		}
		return
	}
}

func respMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		a, err := anypb.New(reply.(proto.Message))
		if err != nil {
			log.Error("响应必须是protobuf类型")
			panic(err)
		}

		var code int
		var msg string

		//var iError errs.IError
		//if ok := errors.As(err, &iError); ok {
		//	code = iError.Code()
		//	msg = iError.Error()
		//}

		reply = &v1.Reply{
			Data: a,
			Code: int32(code),
			Msg:  msg,
		}
		return
	}
}

func UnaryInterceptor() grpc2.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc2.UnaryServerInfo, handler grpc2.UnaryHandler) (resp interface{}, err error) {
		log.Info(req)

		resp, err = handler(ctx, req)

		var code int
		var msg string

		//var iError errs.IError
		//if ok := errors.As(err, &iError); ok {
		//	code = iError.Code()
		//	msg = iError.Error()
		//}

		a, err := anypb.New(resp.(proto.Message))
		if err != nil {
			log.Error(err)
		}
		resp = &v1.Reply{
			Data: a,
			Code: int32(code),
			Msg:  msg,
		}

		return resp, nil
	}
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	se.Code = errs.ErrorReason_value[se.Reason]

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http2.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(http2.StatusOK)
	_, _ = w.Write(body)
}

func ContentType(name string) string {
	switch name {
	case "json":
		return "application/json"
	default:
		return "application/json"
	}
}
