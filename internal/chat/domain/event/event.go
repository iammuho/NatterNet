package event

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/domain/services"
)

const (
	streamName     = "ROOM"
	streamSubjects = "ROOM.*"

	// event types
	MessageCreatedEvent = "ROOM.MESSAGE_CREATED"
)

type roomEvent struct {
	ctx                      context.AppContext
	roomCommandDomainService services.RoomCommandDomainServices
}

func NewRoomEventHandler(ctx context.AppContext, roomCommandDomainService services.RoomCommandDomainServices) {
	// Create the stream
	ctx.GetNatsContext().CreateStream(streamName, streamSubjects)

	r := &roomEvent{
		ctx:                      ctx,
		roomCommandDomainService: roomCommandDomainService,
	}

	// Subscribe to the stream
	ctx.GetNatsContext().Subscribe(streamSubjects, r.handleMessageCreated)
}
