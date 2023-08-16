// Package errorhandler defines the error codes to be used by other packages
package errorhandler

// Error codes
const (
	InternalSystemErrorCode   = 1
	RequestBodyParseErrorCode = 2
	ValidationErrorCode       = 3
)

// Error messages
const (
	RequestBodyParseErrorMessage = "Invalid request body format"
	ValidationErrorMessage       = "Invalid request body"
)
