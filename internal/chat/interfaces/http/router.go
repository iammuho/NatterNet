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

	// Create the chat
	chat := f.Group("/chat", middleware.Protected())
	{
		chat.Post("/room", h.CreateRoom())
	}

}
