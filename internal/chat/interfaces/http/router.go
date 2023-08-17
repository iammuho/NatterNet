package http

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/shared/interfaces/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	ctx context.AppContext
}

// NewChatHandler is the constructor for the chat domain handler
func NewChatHandler(ctx context.AppContext) *handler {
	return &handler{
		ctx: ctx,
	}
}

// RegisterRoutes register the routing handlers to the http server
func (h *handler) RegisterRoutes(f fiber.Router) {
	middleware := http.NewMiddleware(h.ctx)

	// Create the chat
	chat := f.Group("/chat", middleware.Protected())
	{
		chat.Post("/room", h.CreateRoom())
	}

}
