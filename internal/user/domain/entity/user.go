package entity

import (
	"time"

	"github.com/iammuho/natternet/pkg/hashing"
)

type User struct {
	id string

	// Account information
	username string
	password string
	email    string

	// Timestamps
	createdAt time.Time
	updatedAt *time.Time
}

// SetID sets the id
func (u *User) SetID(id string) {
	u.id = id
}

// GetID returns the id
func (u *User) GetID() string {
	return u.id
}

// SetCreatedAt sets the created at
func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

// GetCreatedAt returns the created at
func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

// SetUpdatedAt sets the updated at
func (u *User) SetUpdatedAt(updatedAt *time.Time) {
	u.updatedAt = updatedAt
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

// SetPassword is a function that hashes and saves the password
func (u *User) SetPassword(password string, hash bool) {
	if !hash {
		u.password = password
		return
	}

	hashingFactory := hashing.NewHashingFactory()

	// Hash the password
	pass, _ := hashingFactory.HashPassword(password)

	u.password = pass
}

// GetPassword returns the password
func (u *User) GetPassword() string {
	return u.password
}

// NewUser is a function that creates a new user
func NewUser(uuid string, username string, password string, email string, createdAt time.Time) *User {
	user := &User{}

	// Generate the id
	user.SetID(uuid)

	// Set the username
	user.SetUsername(username)

	// Set the email
	user.SetEmail(email)

	// Generate the timestamps
	user.SetCreatedAt(createdAt)

	// Set the password
	user.SetPassword(password, true)

	return user
}
