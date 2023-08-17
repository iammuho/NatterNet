package dto

type CreateRoomReqDTO struct {
	// Meta
	Name        string `json:"name"`
	Description string `json:"description"`

	// Config
	IsGroup bool `json:"is_group"`

	// Users
	UserIDs []string `json:"user_ids" validate:"required"`
}
