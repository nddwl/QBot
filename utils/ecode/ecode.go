package ecode

import (
	"QBot/utils/config"
	"fmt"
	"strconv"
)

var (
	_codes = map[int]struct{}{}
)

type Codes interface {
	Error() string
	Code() int
	Message() string
}

func New(code int, message string) Code {
	if _, ok := _codes[code]; ok {
		panic("此错误已经存在")
	}
	_codes[code] = struct{}{}
	return Code{
		code:    code,
		message: message,
	}
}

var _ Codes = &Code{}
var _ error = &Code{}

type Code struct {
	code    int
	message string
}

func (e Code) Error() string {
	return "Err" + strconv.FormatInt(int64(e.Code()), 10) + ":" + e.Message()
}

func (e Code) Code() int {
	return e.code
}

func (e Code) Message() string {
	return e.message
}

func (e Code) ReSet(message string) Codes {
	return Code{
		code:    e.code,
		message: message,
	}
}

func Cause(err error) (code Codes) {
	if err == nil {
		return Ok
	}
	if e, ok := err.(Codes); ok {
		return e
	}
	if config.IsLocal() {
		fmt.Printf("\033[1;47;31m未知错误:%s\033\n[0m", err)
	} else {
		fmt.Printf("未知错误:%s\n", err)
	}
	if err.Error() != "" {
		code = Err.ReSet(err.Error())
	} else {
		code = Err
	}
	return
}
