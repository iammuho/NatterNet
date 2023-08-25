package ws

import (
	"sync"

	"github.com/iammuho/natternet/internal/user"

	"github.com/dgrr/websocket"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	application *user.Application
	conn        *websocket.Conn
	clients     sync.Map
	accessToken string
}

// NewWSHandler is the constructor for the ws handler
func NewWSHandler(application *user.Application) *handler {
	return &handler{
		application: application,
		clients:     sync.Map{},
		accessToken: "",
	}
}

// RegisterRoutes register the routing handlers to the ws server
func (h *handler) RegisterRoutes(f fiber.Router) {
	wServer := websocket.Server{}
	wServer.HandleOpen(h.OnOpen)
	wServer.HandleClose(h.OnClose)

	// Create a path for the websocket authenticated connection /ws/:accessToken
	f.Get("/ws/:accessToken", func(c *fiber.Ctx) error {
		h.accessToken = c.Params("accessToken")

		// convert fasthttp ctx to fiber ctx
		ctx := c.Context()
		wServer.Upgrade(ctx)

		return nil
	})
}
