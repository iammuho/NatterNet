// Package errorhandler defines the error codes to be used by other packages
package errorhandler

// Response represents an error response in the system
type Response struct {
	Code       int         `json:"code"`
	Message    interface{} `json:"message"`
	StatusCode int         `json:"status_code"`
}
