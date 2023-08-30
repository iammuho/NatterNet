package dto

type LeaveRoomReqDTO struct {
	RoomID string `validate:"required"`
	UserID string `validate:"required"`
}
