package types

import (
	"encoding/json"
	"time"
)

type WebsocketMessage struct {
	ConnectionID string `json:"connection_id"`

	// Message Type
	MessageType MessageType `json:"type"`

	// Message information
	Message interface{} `json:"message"`

	// Timestamps
	CreatedAt int64 `json:"created_at"`
}

// generateCreatedAt generates a new created at timestamp for the message entity
func (w *WebsocketMessage) generateCreatedAt() {
	w.CreatedAt = time.Now().Unix()
}

// setMessageType sets the message type for the message entity
func (w *WebsocketMessage) setMessageType(messageType MessageType) {
	w.MessageType = messageType
}

// New creates a new message entity
func (w *WebsocketMessage) New(messageType MessageType) {
	w.generateCreatedAt()
	w.setMessageType(messageType)
}

// ToJson converts the message entity to a json byte string
func (w *WebsocketMessage) ToJson() []byte {
	m, jsonErr := json.Marshal(w)
	if jsonErr != nil {
		return nil
	}

	return m
}
