package types

type MessageType string

const (
	MessageTypeError = "error"

	// Health Check
	MessageTypeHealthCheck         = "health.check"
	MessageTypeConnectionConnected = "connection.connected"

	// Message
	MessageTypeMessageCreated = "message.created"

	// Room
	MessageTypeRoomUserJoined = "room.user.joined"
	MessageTypeRoomUserLeft   = "room.user.left"
	MessageTypeRoomEvent      = "room.event"
)
