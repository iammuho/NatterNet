package dto

type QueryUserByIDReqDTO struct {
	UserID string `json:"user_id" validate:"required"`
}
