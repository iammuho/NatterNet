package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Options *HTTPServerOptions
	App     *fiber.App
}

// NewHTTP returns a new HTTP instance
func NewHTTPServer(opts ...Option) *Server {
	options := HTTPServerOptions{}
	for _, o := range opts {
		o(&options)
	}

	config := fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "NatterNet",
		AppName:       "NatterNet App v0.0.1",
		ReadTimeout:   50 * time.Second,
		WriteTimeout:  50 * time.Second,
		BodyLimit:     1024 * 1024 * 1024,
	}

	app := fiber.New(config)

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Return the SERVER
	return &Server{
		Options: &options,
		App:     app,
	}
}
