package response

import "github.com/templatedop/api/modules/server/validation"

type Responser interface {
	ToResponse() any
}

type Stature interface {
	Status() int
}

type Response[T any] struct {
	Success       bool `json:"success"`
	Data          T    `json:"data,omitempty"`
	Page          *int `json:"page,omitempty"`
	Size          *int `json:"size,omitempty"`
	TotalElements *int `json:"totalElements,omitempty"`
	TotalPages    *int `json:"totalPages,omitempty"`

	// Error fields
	Message          string                  `json:"message,omitempty"`
	ValidationErrors []validation.FieldError `json:"validationErrors,omitempty"`
	Errors           []Errors                `json:"errors,omitempty"`
}

type Errors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(data any) Response[any] {
	return Response[any]{
		Success: true,
		Data:    data,
	}
}

func Error(msg string, Errors []Errors, fieldErrs ...validation.FieldError) Response[any] {
	return Response[any]{
		Success:          false,
		Message:          msg,
		ValidationErrors: fieldErrs,
		Errors:           Errors,
	}
}

// type ValidationError struct {
// 	Success          bool                    `json:"success"`
// 	Message          string                  `json:"message,omitempty"`
// 	ValidationErrors []validation.FieldError `json:"Errors,omitempty"`
// }

// func ErrorWithErrors(msg string, errs []Errors) Response[any] {
// 	return Response[any]{
// 		Success: false,
// 		Message: msg,
// 		Errors:  errs,
// 	}
// }
