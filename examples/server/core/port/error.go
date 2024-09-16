package port

import (
	"github.com/templatedop/api/ecode"
)

var (
	CustomCodeDatabaseError = ecode.New(1004, "Database Error", "An error occurred while accessing the database")
	CCodeDatabaseError      = ecode.New(1005, "Database Error", "An error occurred while accessing the database")
	CodeDatabaseError       = ecode.New(1001, "Database Error", "An error occurred while accessing the database.")
	CodeValidationError     = ecode.New(1002, "Validation Error", "Invalid input data.")
)

