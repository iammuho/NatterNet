package event

import (
	"encoding/json"
	"errors"

	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/nats-io/nats.go"
)

func (r *roomEvent) handleMessageCreated(msg *nats.Msg) error {
	msg.Ack()

	// Unmarshal the message
	var event values.MessageValue
	err := json.Unmarshal(msg.Data, &event)

	if err != nil {
		return err
	}

	// Update the last message in the room
	roomErr := r.roomCommandDomainService.UpdateLastMessage(event.RoomID, &event)

	if roomErr != nil {
		return errors.New(roomErr.Message.(string))
	}

	return nil
}
