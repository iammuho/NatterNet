package chat

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/infrastructure/mongodb"
)

// Application represents the application context for chat-related functionality.
type Application struct {
	AppContext         context.AppContext
	RoomCommandHandler application.RoomCommandHandler
	RoomQueryHandler   application.RoomQueryHandler
}

// NewApplication initializes a new chat application context with the given app context.
func NewApplication(ctx context.AppContext) *Application {
	// Setup the room repository
	roomRepository := mongodb.NewRoomRepository(ctx)

	// Setup the domain services
	roomCommandDomainService := services.NewRoomCommandDomainServices(ctx, roomRepository)
	roomQueryDomainService := services.NewRoomQueryDomainServices(ctx, roomRepository)

	// Setup the command handlers
	roomCommandHandler := application.NewRoomCommandHandler(ctx, roomCommandDomainService)
	roomQueryHandler := application.NewRoomQueryHandler(ctx, roomQueryDomainService)

	return &Application{
		AppContext: ctx,

		// Application layer
		RoomCommandHandler: roomCommandHandler,
		RoomQueryHandler:   roomQueryHandler,
	}
}
