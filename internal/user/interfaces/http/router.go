package http

import (
	"github.com/iammuho/natternet/cmd/app/context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	ctx context.AppContext
}

// NewUserHandler is the constructor for the user domain handler
func NewUserHandler(ctx context.AppContext) *handler {
	return &handler{
		ctx: ctx,
	}
}

// RegisterRoutes register the routing handlers to the http server
func (h *handler) RegisterRoutes(f fiber.Router) {
	// Create the user group
	user := f.Group("/user")
	{
		// Create the auth group
		auth := user.Group("/auth")
		{
			auth.Post("/signin", h.Signin)
		}
	}
}

// Signin is the handler for the signin route
func (h *handler) Signin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
