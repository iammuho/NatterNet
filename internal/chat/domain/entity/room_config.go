package entity

type RoomType string

const (
	RoomTypeGroup   RoomType = "group"
	RoomTypePrivate RoomType = "private"
)

type RoomConfig struct {
	RoomType RoomType `json:"room_type" bson:"room_type"`
}

// SetRoomType sets the room type
func (rc *RoomConfig) SetRoomType(roomType RoomType) {
	rc.RoomType = roomType
}

// NewRoomConfig creates a new RoomConfig
// TBD: make this configurable via env vars
func NewRoomConfig(roomType RoomType) *RoomConfig {
	roomConfig := &RoomConfig{}

	roomConfig.SetRoomType(roomType)

	return roomConfig
}
