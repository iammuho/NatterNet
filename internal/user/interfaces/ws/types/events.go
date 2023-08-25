package types

type MessageType string

const (
	MessageTypeError = "error"

	// Health Check
	MessageTypeHealthCheck         = "health.check"
	MessageTypeConnectionConnected = "connection.connected"
)
