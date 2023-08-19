package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/iammuho/natternet/cmd/app/config"
	"github.com/iammuho/natternet/cmd/app/context"
	chatH "github.com/iammuho/natternet/internal/chat/interfaces/http"
	"github.com/iammuho/natternet/internal/user"
	userH "github.com/iammuho/natternet/internal/user/interfaces/http"
	"github.com/iammuho/natternet/pkg/http"
	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// Add the logger
	l, err := logger.NewLogger(
		logger.WithLoggerLevel(config.Config.Logger.Level),
		logger.WithLoggerName(config.Config.Logger.Name),
	)

	if err != nil {
		panic(err)
	}

	// Add the JWT
	jwtContext, err := jwt.NewJWT(
		jwt.WithJWTPublicKeyPath(config.Config.JWT.PublicKeyPath),
		jwt.WithJWTPrivateKeyPath(config.Config.JWT.PrivateKeyPath),
		jwt.WithJWTKeyID(config.Config.JWT.Kid),
		jwt.WithJWTIssuer(config.Config.JWT.Issuer),
		jwt.WithJWTSubject(config.Config.JWT.Subject),
	)

	if err != nil {
		l.Panic("JWT failed to initialize: %v", zap.Error(err))
	}

	// Create a new http server
	l.Info("Creating HTTP Server")
	httpServer := http.NewHTTPServer(
		http.WithHTTPServerHeader(config.Config.Application.Name),
		http.WithHTTPServerAppName(fmt.Sprintf("%s v%s", config.Config.Application.Name, config.Config.Application.Version)),
		http.WithHTTPServerCaseSensitive(config.Config.HTTPServer.CaseSensitive),
		http.WithHTTPServerStrictRouting(config.Config.HTTPServer.StrictRouting),
		http.WithHTTPServerReadTimeout(config.Config.HTTPServer.ReadTimeout),
		http.WithHTTPServerWriteTimeout(config.Config.HTTPServer.WriteTimeout),
		http.WithHTTPServerBodyLimit(config.Config.HTTPServer.BodyLimit),
	)

	// Create the mongo client
	l.Info("Creating MongoDB Client")
	mongodbContext, err := mongodb.NewMongoDB(
		mongodb.WithMongoDBURI(config.Config.MongoDB.URI),
		mongodb.WithMongoDBDatabase(config.Config.MongoDB.Database),
		mongodb.WithMongoDBUsername(config.Config.MongoDB.Username),
		mongodb.WithMongoDBPassword(config.Config.MongoDB.Password),
	)

	if err != nil {
		l.Panic("MongoDB Client failed to connect: %v", zap.Error(err))
	}

	// Create the app context
	ctx := context.NewAppContext(l, jwtContext, mongodbContext)

	// Register the routes
	v1 := httpServer.App.Group("/api/v1")
	{
		// V1 routes
		v1 := v1.Group("/")
		{
			// Create health check route
			v1.Get("/health", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"status": "ok",
				})
			})

			// Setup the user context
			userCtx := user.NewUserApplication(ctx)
			userHandler := userH.NewUserHandler(userCtx)
			userHandler.RegisterRoutes(v1)

			// Chat Handlers
			chatHandler := chatH.NewChatHandler(ctx)
			chatHandler.RegisterRoutes(v1)
		}
	}

	// Start the http server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		l.Info("Starting HTTP Server")
		if err := httpServer.App.Listen(fmt.Sprintf("%s:%d", config.Config.HTTPServer.ListenAddress, config.Config.HTTPServer.ListenPort)); err != nil {
			l.Panic("HTTP Server Listen failed: %v", zap.Error(err))
		}
	}()

	// Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	l.Info("Shutting down the server...")

	if err := httpServer.App.Shutdown(); err != nil {
		l.Panic("HTTP Server Shutdown failed: %v", zap.Error(err))
	}

	wg.Wait()

	l.Info("Server exited")
}
