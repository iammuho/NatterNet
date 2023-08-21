package dto

type CreateRoomReqDTO struct {
	// Meta
	Name        string `json:"name"`
	Description string `json:"description"`

	// Config
	IsGroup bool `json:"is_group"`

	// Users
	Owner   string   `json:"owner" validate:"required"`
	UserIDs []string `json:"user_ids"`
}
