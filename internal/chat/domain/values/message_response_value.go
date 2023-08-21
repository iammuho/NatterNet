package values

import (
	"time"

	"github.com/iammuho/natternet/internal/chat/domain/entity"
)

type MessageValue struct {
	ID string `json:"id" bson:"_id"`

	// Relations
	RoomID   string `json:"room_id" bson:"room_id"`
	SenderID string `json:"sender_id" bson:"sender_id"`

	// Attributes
	Content     string             `json:"content" bson:"content"`
	MessageType entity.MessageType `json:"message_type" bson:"message_type"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

// NewMessageValueFromMessage creates a new message value from message entity
func NewMessageValueFromMessage(message *entity.Message) *MessageValue {
	return &MessageValue{
		ID:          message.GetID(),
		RoomID:      message.GetRoomID(),
		SenderID:    message.GetSenderID(),
		Content:     message.GetContent(),
		MessageType: message.GetMessageType(),
		CreatedAt:   message.GetCreatedAt(),
		UpdatedAt:   message.GetUpdatedAt(),
	}
}
