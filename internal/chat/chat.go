package chat

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/infrastructure/mongodb"
)

// Application represents the application context for chat-related functionality.
type Application struct {
	AppContext            context.AppContext
	RoomCommandHandler    application.RoomCommandHandler
	RoomQueryHandler      application.RoomQueryHandler
	MessageCommandHandler application.MessageCommandHandler
	MessageQueryHandler   application.MessageQueryHandler
}

// NewApplication initializes a new chat application context with the given app context.
func NewApplication(ctx context.AppContext) *Application {
	// Setup the room repository
	roomRepository := mongodb.NewRoomRepository(ctx)
	messageRepository := mongodb.NewMessageRepository(ctx)

	// Setup the domain services
	roomCommandDomainService := services.NewRoomCommandDomainServices(ctx, roomRepository)
	roomQueryDomainService := services.NewRoomQueryDomainServices(ctx, roomRepository)
	messageCommandServices := services.NewMessageCommandDomainServices(ctx, messageRepository, roomQueryDomainService)
	messageQueryServices := services.NewMessageQueryDomainServices(ctx, messageRepository, roomQueryDomainService)

	// Setup the command handlers
	roomCommandHandler := application.NewRoomCommandHandler(ctx, roomCommandDomainService)
	messageCommandHandler := application.NewMessageCommandHandler(ctx, messageCommandServices)

	// Setup the query handlers
	roomQueryHandler := application.NewRoomQueryHandler(ctx, roomQueryDomainService)
	messageQueryHandler := application.NewMessageQueryHandler(ctx, messageQueryServices)

	return &Application{
		AppContext: ctx,

		// Application layer
		RoomCommandHandler:    roomCommandHandler,
		RoomQueryHandler:      roomQueryHandler,
		MessageCommandHandler: messageCommandHandler,
		MessageQueryHandler:   messageQueryHandler,
	}
}
