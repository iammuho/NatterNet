package ws

import (
	"fmt"

	"github.com/iammuho/natternet/internal/user/interfaces/ws/types"
	"go.uber.org/zap"

	"github.com/dgrr/websocket"
)

func (h *handler) OnOpen(c *websocket.Conn) {
	// connection ID to string
	connectionID := fmt.Sprintf("%d", c.ID())

	// Send a connection succesful event
	eventModel := &types.WebsocketMessage{
		ConnectionID: connectionID,
	}
	eventModel.New(types.MessageTypeConnectionConnected)
	eventModel.Message = map[string]string{"message": "Connection successful"}

	// convert eventmodel to []byte
	eventModelByte := eventModel.ToJson()

	if _, err := c.Write(eventModelByte); err != nil {
		h.application.AppContext.GetLogger().Error("Error sending connection successful event",
			zap.String("connectionID", connectionID),
			zap.Error(err))
		return
	}

}

func (h *handler) OnClose(c *websocket.Conn, err error) {
	if err != nil {
		h.application.AppContext.GetLogger().Error("Error closing connection",
			zap.Uint64("connectionID", c.ID()),
			zap.Error(err))
	} else {
		h.application.AppContext.GetLogger().Info("Connection closed", zap.Uint64("connectionID", c.ID()))
	}

	h.clients.Delete(c.ID())
}
