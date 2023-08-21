package dto

type CreateMessageReqDTO struct {
	RoomID      string `validate:"required"`
	SenderID    string `validate:"required"`
	Content     string `json:"content" validate:"required"`
	MessageType string `json:"message_type" validate:"required"`
}
