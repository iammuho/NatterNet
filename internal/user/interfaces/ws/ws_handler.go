package ws

import (
	"fmt"
	"log"

	"github.com/iammuho/natternet/internal/user/interfaces/ws/types"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/dgrr/websocket"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) OnOpen(c *websocket.Conn) {
	// Parse the JWT
	claims, err := h.application.AppContext.GetJwtContext().ParseJWT(h.accessToken)

	if err != nil {
		eventModel := &types.WebsocketMessage{}
		eventModel.New(types.MessageTypeError)
		eventModel.Message = &errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: err, StatusCode: fiber.StatusUnauthorized}

		// convert eventmodel to []byte
		eventModelByte := eventModel.ToJson()

		if _, err := c.Write(eventModelByte); err != nil {
			log.Println(err)
			return
		}

		return
	}

	if claims == nil || claims["ID"] == nil {

		eventModel := &types.WebsocketMessage{}
		eventModel.New(types.MessageTypeError)
		eventModel.Message = &errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: errorhandler.InvalidAccessTokenMessage, StatusCode: fiber.StatusUnauthorized}

		// convert eventmodel to []byte
		eventModelByte := eventModel.ToJson()

		if _, err := c.Write(eventModelByte); err != nil {
			log.Println(err)
			return
		}

		return
	}

	c.SetUserValue("ID", claims["ID"])
	h.clients.Store(c.ID(), c)

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
		log.Println(err)
		return
	}

}

func (h *handler) OnClose(c *websocket.Conn, err error) {
	if err != nil {
		log.Printf("%d closed with error: %s\n", c.ID(), err)
	} else {
		log.Printf("%d closed the connection\n", c.ID())
	}

	h.clients.Delete(c.ID())
}
