package entity

type RoomMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SetName sets the name of the entity
func (r *RoomMeta) SetName(name string) {
	r.Name = name
}

// SetDescription sets the description of the entity
func (r *RoomMeta) SetDescription(description string) {
	r.Description = description
}

// NewRoomMeta creates a new RoomMeta
func NewRoomMeta(name string, description string) *RoomMeta {
	roomMeta := &RoomMeta{}

	roomMeta.SetName(name)
	roomMeta.SetDescription(description)

	return roomMeta
}
