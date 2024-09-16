package errors

import (
	"github.com/templatedop/api/ecode"
)

func (err *Error) Code() ecode.Code {
	if err == nil {
		return ecode.CodeNil
	}
	if err.code == ecode.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

func (err *Error) SetCode(code ecode.Code) {
	if err == nil {
		return
	}
	err.code = code
}
