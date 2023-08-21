package entity

import "time"

type MessageType string

const (
	// text
	MessageTypeText MessageType = "text"
	// image
	MessageTypeImage MessageType = "image"
	// video
	MessageTypeVideo MessageType = "video"
	// audio
	MessageTypeAudio MessageType = "audio"
	// file
	MessageTypeFile MessageType = "file"
	// link
	MessageTypeLink MessageType = "link"
)

// Message is a struct that represents a message in a room
type Message struct {
	id string

	// Relations
	roomID   string
	senderID string

	// Attributes
	content     string
	messageType MessageType

	// Timestamps
	createdAt time.Time
	updatedAt *time.Time
}

// SetID sets the id of the message
func (m *Message) SetID(id string) {
	m.id = id
}

// GetID returns the id of the message
func (m *Message) GetID() string {
	return m.id
}

// SetRoomID sets the room id of the message
func (m *Message) SetRoomID(roomID string) {
	m.roomID = roomID
}

// GetRoomID returns the room id of the message
func (m *Message) GetRoomID() string {
	return m.roomID
}

// SetSenderID sets the sender id of the message
func (m *Message) SetSenderID(senderID string) {
	m.senderID = senderID
}

// GetSenderID returns the sender id of the message
func (m *Message) GetSenderID() string {
	return m.senderID
}

// SetContent sets the content of the message
func (m *Message) SetContent(content string) {
	m.content = content
}

// GetContent returns the content of the message
func (m *Message) GetContent() string {
	return m.content
}

// SetMessageType sets the message type of the message
func (m *Message) SetMessageType(messageType MessageType) {
	m.messageType = messageType
}

// GetMessageType returns the message type of the message
func (m *Message) GetMessageType() MessageType {
	return m.messageType
}

// SetCreatedAt sets the created at timestamp of the message
func (m *Message) SetCreatedAt(createdAt time.Time) {
	m.createdAt = createdAt
}

// GetCreatedAt returns the created at timestamp of the message
func (m *Message) GetCreatedAt() time.Time {
	return m.createdAt
}

// SetUpdatedAt sets the updated at timestamp of the message
func (m *Message) SetUpdatedAt(updatedAt *time.Time) {
	m.updatedAt = updatedAt
}

// GetUpdatedAt returns the updated at timestamp of the message
func (m *Message) GetUpdatedAt() *time.Time {
	return m.updatedAt
}

// NewMessage creates a new message
func NewMessage(
	id string,
	roomID string,
	senderID string,
	content string,
	messageType MessageType,
	createdAt time.Time,
) *Message {
	return &Message{
		id:          id,
		roomID:      roomID,
		senderID:    senderID,
		content:     content,
		messageType: messageType,
		createdAt:   createdAt,
	}
}
