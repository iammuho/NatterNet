package entity

import (
	"time"
)

type Room struct {
	id string

	// Room Details
	meta   RoomMeta
	config RoomConfig
	users  []RoomUser

	// Last Message Fields
	lastMessageID string
	lastMessageAt *time.Time

	// Timestamps
	createdAt time.Time
	updatedAt *time.Time
}

// ID returns the ID of the room entity
func (r *Room) GetID() string {
	return r.id
}

// SetID sets the ID of the room entity
func (r *Room) SetID(id string) {
	r.id = id
}

// SetRoomMeta sets the room meta of the room entity
func (r *Room) SetRoomMeta(meta RoomMeta) {
	r.meta = meta
}

// GetRoomMeta returns the room meta of the room entity
func (r *Room) GetRoomMeta() RoomMeta {
	return r.meta
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
	if !r.CheckRoomUserExists(user.UserID) {
		r.users = append(r.users, user)
	}
}

// CheckRoomUserExists checks if a room user exists in the room entity
func (r *Room) CheckRoomUserExists(userID string) bool {
	for _, user := range r.users {
		if user.UserID == userID {
			return true
		}
	}

	return false
}

// RemoveRoomUser removes a room user from the room entity
func (r *Room) RemoveRoomUser(userID string) {
	for i, user := range r.users {
		if user.UserID == userID {
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

// GetLastMessageID returns the last message ID of the room entity
func (r *Room) GetLastMessageID() string {
	return r.lastMessageID
}

// GetLastMessageAt returns the last message timestamp of the room entity
func (r *Room) GetLastMessageAt() *time.Time {
	return r.lastMessageAt
}

// SetCreatedAt sets the created at timestamp for the room entity
func (r *Room) SetCreatedAt(t time.Time) {
	r.createdAt = t
}

// GetCreatedAt returns the created at timestamp of the room entity
func (r *Room) GetCreatedAt() time.Time {
	return r.createdAt
}

// SetUpdatedAt sets the updated at timestamp for the room entity
func (r *Room) SetUpdatedAt(t *time.Time) {
	r.updatedAt = t
}

// GetUpdatedAt returns the updated at timestamp of the room entity
func (r *Room) GetUpdatedAt() *time.Time {
	return r.updatedAt
}

// NewRoom creates a new room entity
func NewRoom(uuid string, meta RoomMeta, roomType RoomType, createdAt time.Time) *Room {
	room := &Room{}
	room.SetID(uuid)
	room.SetRoomConfig(*NewRoomConfig(roomType))
	room.SetRoomMeta(meta)
	room.SetRoomUsers([]RoomUser{})
	room.SetCreatedAt(createdAt)

	return room
}
