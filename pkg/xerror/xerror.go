// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package xerror

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alimy/mir/v5"
)

var (
	_ mir.Error = (*Error)(nil)

	// codes = map[int]string{}
)

type Error struct {
	code    int
	msg     string
	details []string
}

type ValidError struct {
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func NewError(code int, msg string) *Error {
	// if _, ok := codes[code]; ok {
	// 	panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	// }
	// codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.StatusCode(), e.Msg())
}

func (e *Error) StatusCode() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

func HttpStatusCode(e error) (statusCode int, code int) {
	var ok bool
	statusCode = http.StatusInternalServerError
	if code, ok = mir.HttpStatusCode(e); !ok {
		return
	}
	switch code {
	case Success.StatusCode():
		statusCode = http.StatusOK
	case ServerError.StatusCode():
		statusCode = http.StatusInternalServerError
	case InvalidParams.StatusCode():
		statusCode = http.StatusBadRequest
	case UnauthorizedAuthNotExist.StatusCode(),
		UnauthorizedAuthFailed.StatusCode(),
		UnauthorizedTokenError.StatusCode(),
		UnauthorizedTokenGenerate.StatusCode(),
		UnauthorizedTokenTimeout.StatusCode():
		statusCode = http.StatusUnauthorized
	case TooManyRequests.StatusCode():
		statusCode = http.StatusTooManyRequests
	}
	return
}
