
package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/templatedop/api/ecode"
)


type Error struct {
	error error      
	stack stack      
	text  string     
	code  ecode.Code 
}

const (
	
	stackFilterKeyLocal = "/errors/perror/perror"
)

var (
	
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.ReplaceAll(goRootForFilter, "\\", "/")
	}
}


func (err *Error) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}
	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}


func (err *Error) Cause() error {
	if err == nil {
		return nil
	}
	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				
				loop = e
			} else if e, ok := loop.error.(ICause); ok {
				
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
			
			return errors.New(loop.text)
		}
	}
	return nil
}


func (err *Error) Current() error {
	if err == nil {
		return nil
	}
	return &Error{
		error: nil,
		stack: err.stack,
		text:  err.text,
		code:  err.code,
	}
}


func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.error
}


func (err *Error) Equal(target error) bool {
	if err == target {
		return true
	}
	
	if err.code != Code(target) {
		return false
	}
	
	if err.text != fmt.Sprintf(`%-s`, target) {
		return false
	}
	return true
}
