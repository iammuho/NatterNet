package ws

import (
	"sync"

	"github.com/iammuho/natternet/internal/user"
	eventTypes "github.com/iammuho/natternet/internal/user/domain/event/types"

	"github.com/dgrr/websocket"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	application *user.Application
	clients     sync.Map
	accessToken string
}

// NewWSHandler is the constructor for the ws handler
func NewWSHandler(application *user.Application) *handler {
	// Create the stream
	application.AppContext.GetNatsContext().CreateStream(eventTypes.StreamName, eventTypes.StreamSubjects)

	return &handler{
		application: application,
		clients:     sync.Map{},
	}
}

// RegisterRoutes register the routing handlers to the ws server
func (h *handler) RegisterRoutes(f fiber.Router) {
	wServer := websocket.Server{}
	wServer.HandleOpen(h.OnOpen)
	wServer.HandleClose(h.OnClose)

	// Setup the listeners
	h.setupListeners()

	// Create a path for the websocket authenticated connection /ws/:accessToken
	f.Get("/ws/:accessToken", func(c *fiber.Ctx) error {
		h.accessToken = c.Params("accessToken")

		ctx := c.Context()
		wServer.Upgrade(ctx)

		return nil
	})
}
