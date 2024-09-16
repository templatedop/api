package ecode

import (
	"fmt"

	"sync"
)

var (
	errorCodeRegistry = make(map[int]*Code)
	mu                sync.Mutex
)

type Code interface {
	Code() int

	Message() string

	Detail() interface{}
}

var (
	CodeNil                       = localCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = localCode{0, "OK", nil}                            // It is OK.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // An error occurred internally.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // The given parameter for current operation is invalid.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // Parameter for current operation is missing.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // The function cannot be used like this.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // The configuration is invalid for current operation.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // The configuration is missing for current operation.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // The operation is not implemented yet.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // The operation is not supported yet.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // I tried, but I cannot give you what you want.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Not Authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Security Reason.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // Server is busy, please try again later.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // Resource does not exist.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // Invalid request.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // It needs necessary package import.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // An panic occurred internally.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // Business validation failed.
	CodeMalformedRequest          = localCode{69, "Malformed Request", nil}            // The request is malformed.
	CodeUnMarshallError           = localCode{70, "UnMarshall Error", nil}             // UnMarshall error.
)

func newcode(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}

func New(code int, message string, detail interface{}) Code {
	if code < 1000 {
		panic(fmt.Sprintf("Error code %d must be greater than 1000!", code))
	}
	mu.Lock()
	defer mu.Unlock()

	if _, exists := errorCodeRegistry[code]; exists {
		//fmt.Println("came inside!!")
		panic(fmt.Sprintf("Error code %d already exists!", code))
	}

	newCode := newcode(code, message, detail)
	errorCodeRegistry[code] = &newCode
	return newCode
}
