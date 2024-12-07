package errs

import (
	"fmt"
	"net/http"
	"strconv"
)

type Error struct {
	code int32
	msg  string
}

func (e *Error) Code() int32 {
	return e.code
}

func (e *Error) CodeAsString() string {
	return strconv.FormatInt(int64(e.code), 10)
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d:%s", e.code, e.msg)
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) HttpCode() int {
	if httpCode, ok := errHttp[e.code]; ok {
		return httpCode
	}
	return http.StatusInternalServerError
}

func (e *Error) WithCode(code int32) *Error {
	e.code = code
	return e
}

func (e *Error) WithMessage(msg string) *Error {
	e.msg = msg
	return e
}
