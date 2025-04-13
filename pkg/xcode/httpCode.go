package xcode

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	gozero_errors "github.com/zeromicro/x/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type HttpResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *HttpResponse {
	return &HttpResponse{
		Code: 200,
		Msg: "success",
		Data: data,
	}
}

func Fail(code int, msg string) *HttpResponse {
	return &HttpResponse{
		Code: code,
		Msg: msg,
		Data: nil,
	}
}

func OkHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		code := SERVER_COMMON_ERROR
		msg := CodeToMsg(code)
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*gozero_errors.CodeMsg); ok {
			code = e.Code
			msg = e.Msg
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {				
				code = int(gstatus.Code())
				msg = gstatus.Message()				
			}	
		}
		logx.WithContext(ctx).Errorf("%s error: %v", name, err)
		return http.StatusOK, Fail(code, msg)
	}
}