package websocket

type RoomNewMessageWebsocketValue struct {
	SenderID string      `json:"sender_id"`
	RoomID   string      `json:"room_id"`
	Users    []string    `json:"room_users"`
	Message  interface{} `json:"message"`
}
