package http

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/shared/interfaces/http"

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
	// Create the auth group
	auth := f.Group("/auth")
	{
		auth.Post("/signin", h.Signin())
		auth.Post("/signup", h.Signup())
	}

	middleware := http.NewMiddleware(h.ctx)

	// Create the user me group
	user := f.Group("/user", middleware.Protected())
	{
		user.Get("/me", h.Me())
	}

}
