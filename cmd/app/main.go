package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/iammuho/natternet/cmd/app/config"
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat"
	chatH "github.com/iammuho/natternet/internal/chat/interfaces/http"
	"github.com/iammuho/natternet/internal/user"
	userH "github.com/iammuho/natternet/internal/user/interfaces/http"
	userWS "github.com/iammuho/natternet/internal/user/interfaces/ws"
	"github.com/iammuho/natternet/pkg/http"
	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
	"github.com/iammuho/natternet/pkg/nats"
	"github.com/iammuho/natternet/pkg/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/iammuho/natternet/docs"
	"go.uber.org/zap"
)

// @title NatterNet API
// @version 1.0
// @description NatterNet Chat API Documentation
// @contact.name NatterNet API Support
// @contact.email muhammet.arsln@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1/
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

	// Add the nats
	l.Info("Creating NATS Client")
	natsContext, err := nats.NewNats(
		nats.WithNatsURL(config.Config.Nats.URL),
	)

	if err != nil {
		l.Panic("NATS Client failed to connect: %v", zap.Error(err))
	}

	// Add the storage
	l.Info("Creating Storage", zap.String("driver", config.Config.Storage.Driver))
	storageContext, err := storage.NewStorage(
		storage.WithStorageDriver(config.Config.Storage.Driver),
	)

	if err != nil {
		l.Panic("Storage failed to initialize: %v", zap.Error(err))
	}

	// Create the app context
	ctx := context.NewAppContext(l, jwtContext, mongodbContext, natsContext, storageContext)

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

			v1.Get("/swagger/*", swagger.HandlerDefault)
			v1.Get("/swagger/*", swagger.New(swagger.Config{ // custom
				URL:         "http://example.com/doc.json",
				DeepLinking: false,
				// Expand ("list") or Collapse ("none") tag groups by default
				DocExpansion: "none",
				// Prefill OAuth ClientId on Authorize popup
				OAuth: &swagger.OAuthConfig{
					AppName:  "OAuth Provider",
					ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
				},
				// Ability to change OAuth2 redirect uri location
				OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
			}))

			// Setup the user context
			userApp := user.NewApplication(ctx)

			// User HTTP
			userHandler := userH.NewUserHandler(userApp)
			userHandler.RegisterRoutes(v1)

			// User WS
			userWSHandler := userWS.NewWSHandler(userApp)
			userWSHandler.RegisterRoutes(v1)

			// Chat Handlers
			chatApp := chat.NewApplication(ctx)
			chatHandler := chatH.NewChatHandler(chatApp)
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
