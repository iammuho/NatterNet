package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/iammuho/natternet/internal/user/infrastructure/hashing"
)

type User struct {
	id string `json:"id" bson:"_id"`

	// Account information
	username string `json:"username" bson:"username"`
	password string `json:"password" bson:"password"`
	email    string `json:"email" bson:"email"`

	// Timestamps
	createdAt time.Time  `json:"created_at" bson:"created_at"`
	updatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

func (u *User) generateID() {
	u.id = uuid.New().String()
}

// GetID returns the id
func (u *User) GetID() string {
	return u.id
}

func (u *User) generateCreatedAt() {
	u.createdAt = time.Now()
}

// GetCreatedAt returns the created at
func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) generateUpdatedAt() {
	now := time.Now()
	u.updatedAt = &now
}

// GetUpdatedAt returns the updated at
func (u *User) GetUpdatedAt() *time.Time {
	return u.updatedAt
}

// SetEmail sets the email
func (u *User) SetEmail(email string) {
	u.email = email
}

// GetEmail returns the email
func (u *User) GetEmail() string {
	return u.email
}

// SetUsername sets the username
func (u *User) SetUsername(username string) {
	u.username = username
}

// GetUsername returns the username
func (u *User) GetUsername() string {
	return u.username
}

// setPassword is a function that hashes and saves the password
func (u *User) SetPassword(password string) {
	hashingFactory := hashing.NewHashingFactory()

	// Hash the password
	pass, _ := hashingFactory.HashPassword(password)

	u.password = pass
}

func (u *User) ComparePassword(password string) bool {
	hashingFactory := hashing.NewHashingFactory()

	// Compare the password
	return hashingFactory.ComparePassword(password, u.password)
}

// NewUser is a function that creates a new user
func NewUser(username string, password string, email string) *User {
	user := &User{}

	// Generate the id
	user.generateID()

	// Set the username
	user.SetUsername(username)

	// Set the email
	user.SetEmail(email)

	// Generate the timestamps
	user.generateCreatedAt()

	// Set the password
	user.SetPassword(password)

	return user
}
