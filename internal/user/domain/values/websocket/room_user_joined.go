package websocket

type RoomUserJoinedWebsocketValue struct {
	Username string   `json:"username"`
	UserID   string   `json:"user_id"`
	Users    []string `json:"room_users"`
	RoomID   string   `json:"room_id"`
}
