package errors

import (
	//"fmt"
	"strings"
	//"sync"

	"github.com/templatedop/api/ecode"
)


// Has to be removed
func new(text string) error {
	return &Error{
		stack: callers(),
		text:  text,
		code:  ecode.CodeNil,
	}
}

func NewCode(code ecode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

func WrapCode(code ecode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  text,
		code:  Code(err),
	}
}

func Code(err error) ecode.Code {
	if err == nil {
		return ecode.CodeNil
	}
	if e, ok := err.(ICode); ok {
		return e.Code()
	}
	if e, ok := err.(IUnwrap); ok {
		return Code(e.Unwrap())
	}
	return ecode.CodeNil
}

func HasCode(err error, code ecode.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(ICode); ok && code == e.Code() {
		return true
	}
	if e, ok := err.(IUnwrap); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
