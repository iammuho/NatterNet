package http

import (
	"github.com/iammuho/natternet/internal/chat"
	"github.com/iammuho/natternet/internal/shared/interfaces/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	application *chat.Application
}

// NewChatHandler is the constructor for the chat domain handler
func NewChatHandler(application *chat.Application) *handler {
	return &handler{
		application: application,
	}
}

// RegisterRoutes register the routing handlers to the http server
func (h *handler) RegisterRoutes(f fiber.Router) {
	middleware := http.NewMiddleware(h.application.AppContext)

	chat := f.Group("/chat", middleware.Protected())
	{
		// Room routes
		chat.Post("/room", h.createRoom())
		chat.Get("/rooms", h.queryRooms())
		chat.Post("/room/:roomID/join", h.joinRoom())
		chat.Post("/room/:roomID/leave", h.leaveRoom())
		chat.Post("/room/:roomID/event", h.sendRoomEvent())

		// Message routes
		chat.Post("/room/:roomID/message", h.createMessage())
		chat.Get("/room/:roomID/messages", h.queryMessages())
	}

}
