package xcode

import (
	"github.com/zeromicro/x/errors"
)

func NewError(code int, msg string) error {
	return errors.New(code, msg)
}

func NewDatabaseError() error {
	return errors.New(DATABASE_ERROR, CodeToMsg(DATABASE_ERROR))
}

func NewRequestError() error {
	return errors.New(REQUEST_PARAM_ERROR, CodeToMsg(REQUEST_PARAM_ERROR))
}

func NewServerCommonError() error {
	return errors.New(SERVER_COMMON_ERROR, CodeToMsg(SERVER_COMMON_ERROR))
}
