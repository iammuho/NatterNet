package dto

type CreateMessageReqDTO struct {
	RoomID      string `swaggerignore:"true" validate:"required"`
	SenderID    string `swaggerignore:"true" validate:"required"`
	Content     string `json:"content" validate:"required"`
	MessageType string `json:"message_type" validate:"required"`
}
