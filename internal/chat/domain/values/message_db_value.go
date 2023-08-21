package values

import (
	"time"

	"github.com/iammuho/natternet/internal/chat/domain/entity"
)

type MessageDBValue struct {
	ID string `bson:"_id"`

	// Relations
	RoomID   string `bson:"room_id"`
	SenderID string `bson:"sender_id"`

	// Attributes
	Content     string             `bson:"content"`
	MessageType entity.MessageType `bson:"message_type"`

	// Timestamps
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}

// NewMessageDBValueFromMessage creates a new message db value from message entity
func NewMessageDBValueFromMessage(message *entity.Message) *MessageDBValue {
	return &MessageDBValue{
		ID:          message.GetID(),
		RoomID:      message.GetRoomID(),
		SenderID:    message.GetSenderID(),
		Content:     message.GetContent(),
		MessageType: message.GetMessageType(),
		CreatedAt:   message.GetCreatedAt(),
		UpdatedAt:   message.GetUpdatedAt(),
	}
}

// ToMessage converts the message db value to message entity
func (m MessageDBValue) ToMessage() *entity.Message {
	message := entity.NewMessage(
		m.ID,
		m.RoomID,
		m.SenderID,
		m.Content,
		m.MessageType,
		m.CreatedAt,
	)

	return message
}
