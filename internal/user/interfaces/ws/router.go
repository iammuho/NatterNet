package ws

import (
	"sync"

	"github.com/iammuho/natternet/internal/user"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/dgrr/websocket"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	application *user.Application
	clients     sync.Map
}

// NewWSHandler is the constructor for the ws handler
func NewWSHandler(application *user.Application) *handler {
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

	// Create a path for the websocket authenticated connection /ws/:accessToken
	f.Get("/ws/:accessToken", func(c *fiber.Ctx) error {
		accessToken := c.Params("accessToken")

		claims, err := h.application.AppContext.GetJwtContext().ParseJWT(accessToken)
		if err != nil {
			return h.sendErrorResponse(c, errorhandler.InvalidAccessTokenErrorCode, err.Message.(string))
		}

		if claims == nil || claims["ID"] == nil {
			return h.sendErrorResponse(c, errorhandler.InvalidAccessTokenErrorCode, errorhandler.InvalidAccessTokenMessage)
		}

		ctx := c.Context()
		wServer.Upgrade(ctx)

		ctx.SetUserValue("ID", claims["ID"])
		h.clients.Store(ctx.ID(), c)

		return nil
	})
}

func (h *handler) sendErrorResponse(c *fiber.Ctx, code int, message string) error {
	response := &errorhandler.Response{Code: code, Message: message, StatusCode: fiber.StatusUnauthorized}
	return c.Status(fiber.StatusUnauthorized).JSON(response)
}
