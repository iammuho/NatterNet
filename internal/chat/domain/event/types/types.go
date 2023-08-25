package types

const (
	StreamName     = "ROOM"
	StreamSubjects = "ROOM.*"
)

// Event Types
const RoomCreatedEvent = "ROOM.CREATED"
const RoomUpdatedEvent = "ROOM.UPDATED"
const RoomDeletedEvent = "ROOM.DELETED"
const MessageCreatedEvent = "ROOM.MESSAGE_CREATED"
