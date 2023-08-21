package values

import (
	"time"

	"github.com/iammuho/natternet/internal/chat/domain/entity"
)

type RoomValue struct {
	ID string `json:"id" bson:"_id"`

	// Room Details
	Meta   entity.RoomMeta   `json:"meta" bson:"meta"`
	Config entity.RoomConfig `json:"config" bson:"config"`
	Users  []entity.RoomUser `json:"users" bson:"users"`

	// Last Message Fields
	LastMessageID string     `json:"last_message_id" bson:"last_message_id"`
	LastMessageAt *time.Time `json:"last_message_at" bson:"last_message_at"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

// NewRoomValueFromRoom converts a room entity to a room value
func NewRoomValueFromRoom(room *entity.Room) *RoomValue {
	return &RoomValue{
		ID:        room.GetID(),
		Meta:      room.GetRoomMeta(),
		Config:    room.GetRoomConfig(),
		Users:     room.GetRoomUsers(),
		CreatedAt: room.GetCreatedAt(),
		UpdatedAt: room.GetUpdatedAt(),
	}
}

// ToRoom converts a room value to a room entity
func (r *RoomValue) ToRoom() *entity.Room {
	room := &entity.Room{}

	room.SetID(r.ID)
	room.SetRoomMeta(r.Meta)
	room.SetRoomConfig(r.Config)
	room.SetRoomUsers(r.Users)
	room.SetCreatedAt(r.CreatedAt)
	room.SetUpdatedAt(r.UpdatedAt)

	return room
}
