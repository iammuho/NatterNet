package http

import (
	"github.com/iammuho/natternet/internal/shared/interfaces/http"
	"github.com/iammuho/natternet/internal/user"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	application *user.Application
}

// NewUserHandler is the constructor for the user domain handler
func NewUserHandler(application *user.Application) *handler {
	return &handler{
		application: application,
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

	middleware := http.NewMiddleware(h.application.AppContext)

	// Create the user me group
	user := f.Group("/user", middleware.Protected())
	{
		user.Get("/me", h.Me())
	}

}
