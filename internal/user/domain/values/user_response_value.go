package values

import (
	"time"

	"github.com/iammuho/natternet/internal/user/domain/entity"
)

type UserValue struct {
	ID string `json:"id" bson:"_id"`

	// Account information
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

// NewUserValueFromUser converts a user entity to a user value
func NewUserValueFromUser(user *entity.User) *UserValue {
	return &UserValue{
		ID:        user.GetID(),
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}
