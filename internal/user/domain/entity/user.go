package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/iammuho/natternet/internal/user/infrastructure/hashing"
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

func (u *User) generateID() {
	u.id = uuid.New().String()
}

// SetID sets the id
func (u *User) SetID(id string) {
	u.id = id
}

// GetID returns the id
func (u *User) GetID() string {
	return u.id
}

func (u *User) generateCreatedAt() {
	u.createdAt = time.Now()
}

// SetCreatedAt sets the created at
func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

// GetCreatedAt returns the created at
func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) generateUpdatedAt() {
	now := time.Now()
	u.updatedAt = &now
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

// ComparePassword is a function that compares the password
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
	user.SetPassword(password, true)

	return user
}
