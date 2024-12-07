package errs

import (
	"net/http"
)

var (
	Success       = New(0, "success")
	InternalError = New(100001, "internal error")
	ParamError    = New(100002, "param error")
	AuthError     = New(100011, "auth error")
	NotFoundError = New(100012, "not found")
)

var (
	errHttp = map[int32]int{
		Success.Code():       http.StatusOK,
		ParamError.Code():    http.StatusBadRequest,
		InternalError.Code(): http.StatusInternalServerError,
	}
)

func New(code int32, msg string) *Error {
	return &Error{code, msg}
}
