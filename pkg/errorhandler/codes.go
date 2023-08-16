// Package errorhandler defines the error codes to be used by other packages
package errorhandler

// Error codes
const (
	InternalSystemErrorCode   = 1
	RequestBodyParseErrorCode = 2
	ValidationErrorCode       = 3
	DatabaseErrorCode         = 4

	// Auth related error codes
	InvalidCredentialsErrorCode = 100
)

// Error messages
const (
	RequestBodyParseErrorMessage = "Invalid request body format"
	ValidationErrorMessage       = "Invalid request body"

	// Auth related error messages
	InvalidCredentialsMessage = "Invalid credentials"
)
