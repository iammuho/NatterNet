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
	r := &roomEvent{
		ctx:                      ctx,
		roomCommandDomainService: roomCommandDomainService,
	}

	// Create the stream
	ctx.GetNatsContext().CreateStream(streamName, streamSubjects)

	// Subscribe to the stream
	ctx.GetNatsContext().Subscribe(streamSubjects, r.handleMessageCreated)
}
