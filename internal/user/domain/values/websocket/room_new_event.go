package websocket

type RoomNewEventWebsocketValue struct {
	RoomID    string   `json:"room_id"`
	SenderID  string   `json:"sender_id"`
	UserIDs   []string `json:"user_ids"`
	EventType string   `json:"event"`
}
