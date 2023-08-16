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
	userID string
	role   RoomUserRole

	// Status
	status RoomUserStatus

	// Booleans
	isMuted bool

	// Timestamps
	createdAt time.Time
	updatedAt *time.Time
}

// SetUserID sets the user id
func (r *RoomUser) SetUserID(userID string) {
	r.userID = userID
}

// GetUserID returns the user id
func (r *RoomUser) GetUserID() string {
	return r.userID
}

// SetRole sets the role
func (r *RoomUser) SetRole(role RoomUserRole) {
	r.role = role
}

// GetRole returns the role
func (r *RoomUser) GetRole() RoomUserRole {
	return r.role
}

// SetStatus sets the status
func (r *RoomUser) SetStatus(status RoomUserStatus) {
	r.status = status
}

// GetStatus returns the status
func (r *RoomUser) GetStatus() RoomUserStatus {
	return r.status
}

// SetIsMuted sets the is muted
func (r *RoomUser) SetIsMuted(isMuted bool) {
	r.isMuted = isMuted
}

// GetIsMuted returns the is muted
func (r *RoomUser) GetIsMuted() bool {
	return r.isMuted
}

// SetCreatedAt sets the created at
func (r *RoomUser) SetCreatedAt(createdAt time.Time) {
	r.createdAt = createdAt
}

// generateCreatedAt generates a new created at timestamp for the room entity
func (r *RoomUser) generateCreatedAt() {
	r.createdAt = time.Now()
}

// GetCreatedAt returns the created at
func (r *RoomUser) GetCreatedAt() time.Time {
	return r.createdAt
}

// SetUpdatedAt sets the updated at
func (r *RoomUser) SetUpdatedAt(updatedAt time.Time) {
	r.updatedAt = &updatedAt
}

// GetUpdatedAt returns the updated at
func (r *RoomUser) GetUpdatedAt() *time.Time {
	return r.updatedAt
}

// generateUpdatedAt generates a new updated at timestamp for the room entity
func (r *RoomUser) generateUpdatedAt() {
	now := time.Now()
	r.updatedAt = &now
}

// NewRoomUser creates a new room user
func NewRoomUser(userID string, role RoomUserRole) *RoomUser {
	roomUser := &RoomUser{}

	roomUser.SetUserID(userID)
	roomUser.SetRole(role)
	roomUser.SetStatus(RoomUserStatusActive)
	roomUser.SetIsMuted(false)
	roomUser.generateCreatedAt()

	return roomUser
}
