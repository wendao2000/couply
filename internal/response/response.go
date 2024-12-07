package response

import (
	"github.com/wendao2000/couply/pkg/errs"
)

type BaseResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func New(code, msg string) *BaseResp {
	return &BaseResp{
		Code:    code,
		Message: msg,
	}
}

func Success() *BaseResp {
	return &BaseResp{
		Code:    errs.Success.CodeAsString(),
		Message: errs.Success.Message(),
	}
}

func FromError(err *errs.Error) *BaseResp {
	if err == nil {
		return Success()
	}
	return &BaseResp{
		Code:    err.CodeAsString(),
		Message: err.Message(),
	}
}
