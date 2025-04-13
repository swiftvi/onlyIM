package rpcserver

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	gozero_errors "github.com/zeromicro/x/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err == nil {
		return resp, nil
	}
	logx.WithContext(ctx).Errorf("Rpc service error: %v", err)

	causeErr := errors.Cause(err)
	//类型断言带检测（Comma-ok 断言）
	if e, ok := causeErr.(*gozero_errors.CodeMsg); ok {
		err = status.Error(codes.Code(e.Code), e.Msg)
	}
	return resp, err
}