package entity

type RoomType string

const (
	RoomTypeGroup   RoomType = "group"
	RoomTypePrivate RoomType = "private"
)

type RoomConfig struct {
	maxUsers int
	roomType RoomType
}

// SetMaxUsers sets the max number of users in a room
func (rc *RoomConfig) SetMaxUsers(maxUsers int) {
	rc.maxUsers = maxUsers
}

// GetMaxUsers gets the max number of users in a room
func (rc *RoomConfig) GetMaxUsers() int {
	return rc.maxUsers
}

// SetRoomType sets the room type
func (rc *RoomConfig) SetRoomType(roomType RoomType) {
	rc.roomType = roomType
}

// GetRoomType gets the room type
func (rc *RoomConfig) GetRoomType() RoomType {
	return rc.roomType
}

// NewRoomConfig creates a new RoomConfig
// TBD: make this configurable via env vars
func NewRoomConfig() *RoomConfig {
	roomConfig := &RoomConfig{}

	roomConfig.SetMaxUsers(2)
	roomConfig.SetRoomType(RoomTypePrivate)

	return roomConfig
}
