package dto

type JoinRoomReqDTO struct {
	RoomID string `validate:"required"`
	UserID string `validate:"required"`
}
