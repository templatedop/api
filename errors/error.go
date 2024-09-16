
package errors

import (
	"github.com/templatedop/api/ecode"
)


type IEqual interface {
	Error() string
	Equal(target error) bool
}


type ICode interface {
	Error() string
	Code() ecode.Code
}


type IStack interface {
	Error() string
	Stack() string
}


type ICause interface {
	Error() string
	Cause() error
}


type ICurrent interface {
	Error() string
	Current() error
}


type IUnwrap interface {
	Error() string
	Unwrap() error
}

const (
	
	commaSeparatorSpace = ", "
)
