package values

import (
	"time"

	"github.com/iammuho/natternet/internal/chat/domain/entity"
)

type RoomDBValue struct {
	ID string `bson:"_id"`

	// Room Details
	Meta   entity.RoomMeta   `bson:"meta"`
	Config entity.RoomConfig `bson:"config"`
	Users  []entity.RoomUser `bson:"users"`

	// Last Message Fields
	LastMessageID string     `bson:"last_message_id"`
	LastMessageAt *time.Time `bson:"last_message_at"`

	// Timestamps
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}

func (r RoomDBValue) ToRoom() entity.Room {
	room := entity.Room{}

	room.SetID(r.ID)
	room.SetRoomMeta(r.Meta)
	room.SetRoomConfig(r.Config)
	room.SetRoomUsers(r.Users)
	room.SetLastMessage(r.LastMessageID)
	room.SetCreatedAt(r.CreatedAt)
	room.SetUpdatedAt(r.UpdatedAt)

	return room
}

func NewRoomDBValue(room *entity.Room) *RoomDBValue {
	return &RoomDBValue{
		ID:            room.GetID(),
		Meta:          room.GetRoomMeta(),
		Config:        room.GetRoomConfig(),
		Users:         room.GetRoomUsers(),
		LastMessageID: room.GetLastMessageID(),
		LastMessageAt: room.GetLastMessageAt(),
		CreatedAt:     room.GetCreatedAt(),
		UpdatedAt:     room.GetUpdatedAt(),
	}
}
