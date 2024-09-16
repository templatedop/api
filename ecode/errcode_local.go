package ecode

//import "fmt"

type localCode struct {
	code    int
	message string
	detail  interface{}
}

func (c localCode) Code() int {
	return c.code
}

func (c localCode) Message() string {
	return c.message
}

func (c localCode) Detail() interface{} {
	return c.detail
}

