package values

import (
	"time"

	"github.com/iammuho/natternet/internal/user/domain/entity"
)

type UserDBValue struct {
	ID string `bson:"_id"`

	// Account information
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`

	// Timestamps
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}

// NewUserDBValueFromUser converts a user entity to a user value
func NewUserDBValueFromUser(user *entity.User) *UserDBValue {
	return &UserDBValue{
		ID:        user.GetID(),
		Username:  user.GetUsername(),
		Password:  user.GetPassword(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}

// ToUserEntity converts a user db value to a user entity
func (u *UserDBValue) ToUserEntity() *entity.User {
	user := &entity.User{}

	user.SetID(u.ID)
	user.SetUsername(u.Username)
	user.SetPassword(u.Password, false)
	user.SetEmail(u.Email)
	user.SetCreatedAt(u.CreatedAt)
	user.SetUpdatedAt(u.UpdatedAt)

	return user
}
