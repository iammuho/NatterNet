package entity

type RoomType string

const (
	RoomTypeGroup   RoomType = "group"
	RoomTypePrivate RoomType = "private"
)

type RoomConfig struct {
	MaxUsers int      `json:"max_users" bson:"max_users"`
	RoomType RoomType `json:"room_type" bson:"room_type"`
}

// SetMaxUsers sets the max number of users in a room
func (rc *RoomConfig) SetMaxUsers(maxUsers int) {
	rc.MaxUsers = maxUsers
}

// SetRoomType sets the room type
func (rc *RoomConfig) SetRoomType(roomType RoomType) {
	rc.RoomType = roomType
}

// NewRoomConfig creates a new RoomConfig
// TBD: make this configurable via env vars
func NewRoomConfig(roomType RoomType) *RoomConfig {
	roomConfig := &RoomConfig{}

	roomConfig.SetMaxUsers(2)
	roomConfig.SetRoomType(roomType)

	return roomConfig
}
