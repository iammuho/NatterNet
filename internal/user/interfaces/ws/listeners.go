package ws

import (
	"encoding/json"
	"fmt"

	"github.com/dgrr/websocket"
	eventTypes "github.com/iammuho/natternet/internal/user/domain/event/types"
	websocketValues "github.com/iammuho/natternet/internal/user/domain/values/websocket"
	"github.com/iammuho/natternet/internal/user/interfaces/ws/types"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// setup listeners
func (h *handler) setupListeners() {
	h.application.AppContext.GetNatsContext().Subscribe(eventTypes.StreamSubjects, func(msg *nats.Msg) error {
		msg.Ack()
		h.application.AppContext.GetLogger().Logger.Info("Received message on subject: ", zap.String("subject", msg.Subject))
		switch msg.Subject {
		case eventTypes.MessageCreatedEvent:
			return h.onMessageCreated(msg)
		case eventTypes.RoomUserJoinedEvent:
			return h.onUserJoinedRoom(msg)
		}

		return nil
	})
}

// onMessageCreated handles the message created event
func (h *handler) onMessageCreated(msg *nats.Msg) error {
	msg.Ack()

	// Unmarshal the message
	var event websocketValues.RoomNewMessageWebsocketValue
	err := json.Unmarshal(msg.Data, &event)

	if err != nil {
		return err
	}

	// Check if the user ID is in clients
	h.clients.Range(func(client, v interface{}) bool {
		nc := v.(*websocket.Conn)

		// Range the users
		for _, user := range event.Users {

			if nc.UserValue("ID").(string) == user {
				h.application.AppContext.GetLogger().Logger.Info("Sending message created to ws client with ID: ", zap.String("ID", nc.UserValue("ID").(string)))

				// Create the event model
				eventModel := &types.WebsocketMessage{}
				eventModel.New(types.MessageTypeMessageCreated)
				eventModel.ConnectionID = fmt.Sprintf("%d", nc.ID())
				eventModel.Message = event.Message

				nc.Write(eventModel.ToJson())

				return true
			}
		}

		return true
	})

	return nil
}

// onUserJoinedRoom handles the user joined room event
func (h *handler) onUserJoinedRoom(msg *nats.Msg) error {
	msg.Ack()

	// Unmarshal the message
	var event websocketValues.RoomUserJoinedWebsocketValue
	err := json.Unmarshal(msg.Data, &event)

	if err != nil {
		return err
	}

	// Check if the user ID is in clients
	h.clients.Range(func(client, v interface{}) bool {
		nc := v.(*websocket.Conn)

		// If the user ID is the same as the client ID, skip
		if nc.UserValue("ID").(string) == event.UserID {
			return true
		}

		// Range the users
		for _, user := range event.Users {
			// If the user ID is the same as the client ID, skip
			if nc.UserValue("ID").(string) == user {
				h.application.AppContext.GetLogger().Logger.Info("Sending room user joined to ws client with ID: ", zap.String("ID", nc.UserValue("ID").(string)))

				// Create the event model
				eventModel := &types.WebsocketMessage{}
				eventModel.New(types.MessageTypeRoomUserJoined)
				eventModel.ConnectionID = fmt.Sprintf("%d", nc.ID())
				eventModel.Message = event

				nc.Write(eventModel.ToJson())

				return true

			}

		}

		return true
	})

	return nil
}
