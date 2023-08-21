// Package errorhandler defines the error codes to be used by other packages
package errorhandler

// Error codes
const (
	InternalSystemErrorCode   = 1
	RequestBodyParseErrorCode = 2
	ValidationErrorCode       = 3
	DatabaseErrorCode         = 4

	// Auth related error codes
	InvalidCredentialsErrorCode    = 100
	EmailAlreadyExistsErrorCode    = 101
	UsernameAlreadyExistsErrorCode = 102
	InvalidAccessTokenErrorCode    = 103
	InvalidRefreshTokenErrorCode   = 104
	ExpiredAccessTokenErrorCode    = 105
	ExpiredRefreshTokenErrorCode   = 106

	// User related error codes
	UserNotFoundErrorCode = 200

	// Chat related error codes
	RoomNotFoundErrorCode = 300
	UserIsNotInRoomCode   = 301
)

// Error messages
const (
	RequestBodyParseErrorMessage = "Invalid request body format"
	ValidationErrorMessage       = "Invalid request body"

	// Auth related error messages
	InvalidCredentialsMessage    = "Invalid credentials"
	EmailAlreadyExistsMessage    = "Email already exists"
	UsernameAlreadyExistsMessage = "Username already exists"
	InvalidAccessTokenMessage    = "Invalid access token"
	InvalidRefreshTokenMessage   = "Invalid refresh token"
	ExpiredAccessTokenMessage    = "Access token expired"
	ExpiredRefreshTokenMessage   = "Refresh token expired"

	// User related error messages
	UserNotFoundMessage = "User not found"

	// Chat related error messages
	RoomNotFoundMessage    = "Room not found"
	UserIsNotInRoomMessage = "User is not in room"
)
