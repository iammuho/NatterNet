package entity

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	id string

	// Room Details
	config RoomConfig
	users  []RoomUser

	// Last Message Fields
	lastMessageID string
	lastMessageAt *time.Time

	// Timestamps
	createdAt time.Time
	updatedAt *time.Time
}

// generateID generates a new ID for the room entity
func (r *Room) generateID() {
	r.id = uuid.New().String()
}

// ID returns the ID of the room entity
func (r *Room) GetID() string {
	return r.id
}

// SetID sets the ID of the room entity
func (r *Room) SetID(id string) {
	r.id = id
}

// SetRoomConfig sets the room config of the room entity
func (r *Room) SetRoomConfig(config RoomConfig) {
	r.config = config
}

// GetRoomConfig returns the room config of the room entity
func (r *Room) GetRoomConfig() RoomConfig {
	return r.config
}

// SetRoomUsers sets the room users of the room entity
func (r *Room) SetRoomUsers(users []RoomUser) {
	r.users = users
}

// GetRoomUsers returns the room users of the room entity
func (r *Room) GetRoomUsers() []RoomUser {
	return r.users
}

// AddRoomUser adds a room user to the room entity
func (r *Room) AddRoomUser(user RoomUser) {
	if !r.CheckRoomUserExists(user.GetUserID()) {
		r.users = append(r.users, user)
	}
}

// CheckRoomUserExists checks if a room user exists in the room entity
func (r *Room) CheckRoomUserExists(userID string) bool {
	for _, user := range r.users {
		if user.GetUserID() == userID {
			return true
		}
	}

	return false
}

// RemoveRoomUser removes a room user from the room entity
func (r *Room) RemoveRoomUser(userID string) {
	for i, user := range r.users {
		if user.GetUserID() == userID {
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}
}

// SetLastMessage sets the last message of the room entity
func (r *Room) SetLastMessage(id string) {
	now := time.Now()
	r.lastMessageID = id
	r.lastMessageAt = &now
}

// generateCreatedAt generates a new created at timestamp for the room entity
func (r *Room) generateCreatedAt() {
	r.createdAt = time.Now()
}

// generateUpdatedAt generates a new updated at timestamp for the room entity
func (r *Room) generateUpdatedAt() {
	now := time.Now()
	r.updatedAt = &now
}

// NewRoom creates a new room entity
func NewRoom(config RoomConfig, users []RoomUser) *Room {
	room := &Room{}
	room.generateID()
	room.SetRoomConfig(config)
	room.SetRoomUsers(users)
	room.generateCreatedAt()
	room.generateUpdatedAt()

	return room
}
