package entity

import "time"

type RoomUserRole string

const (
	RoomUserRoleAdmin  RoomUserRole = "admin"
	RoomUserRoleMember RoomUserRole = "member"
)

type RoomUserStatus string

const (
	RoomUserStatusActive RoomUserStatus = "active"
	RoomUserStatusLeft   RoomUserStatus = "left"
)

type RoomUser struct {
	// User Details
	UserID string       `json:"user_id" bson:"user_id"`
	Role   RoomUserRole `json:"role" bson:"role"`

	// Status
	Status RoomUserStatus `json:"status" bson:"status"`

	// Booleans
	IsMuted bool `json:"is_muted" bson:"is_muted"`

	// Timestamps
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}

// SetUserID sets the user id
func (r *RoomUser) SetUserID(userID string) {
	r.UserID = userID
}

// SetRole sets the role
func (r *RoomUser) SetRole(role RoomUserRole) {
	r.Role = role
}

// SetStatus sets the status
func (r *RoomUser) SetStatus(status RoomUserStatus) {
	r.Status = status
}

// SetIsMuted sets the is muted
func (r *RoomUser) SetIsMuted(isMuted bool) {
	r.IsMuted = isMuted
}

// SetCreatedAt sets the created at
func (r *RoomUser) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

// generateCreatedAt generates a new created at timestamp for the room entity
func (r *RoomUser) generateCreatedAt() {
	r.CreatedAt = time.Now()
}

// SetUpdatedAt sets the updated at
func (r *RoomUser) SetUpdatedAt(updatedAt time.Time) {
	r.UpdatedAt = &updatedAt
}

// generateUpdatedAt generates a new updated at timestamp for the room entity
func (r *RoomUser) generateUpdatedAt() {
	now := time.Now()
	r.UpdatedAt = &now
}

// NewRoomUser creates a new room user
func NewRoomUser(userID string, role RoomUserRole, createdAt time.Time) *RoomUser {
	roomUser := &RoomUser{}

	roomUser.SetUserID(userID)
	roomUser.SetRole(role)
	roomUser.SetStatus(RoomUserStatusActive)
	roomUser.SetIsMuted(false)
	roomUser.SetCreatedAt(createdAt)

	return roomUser
}
