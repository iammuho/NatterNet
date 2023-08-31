package dto

// SendRoomEventReqDTO is the request DTO for sending room events
type SendRoomEventReqDTO struct {
	RoomID    string `validate:"required"`
	UserID    string `validate:"required"`
	EventType string `json:"event_type" validate:"required"`
}
