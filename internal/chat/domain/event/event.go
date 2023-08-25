package event

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/domain/event/types"
	"github.com/iammuho/natternet/internal/chat/domain/services"
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
	ctx.GetNatsContext().CreateStream(types.StreamName, types.StreamSubjects)

	// Subscribe to the stream
	ctx.GetNatsContext().Subscribe(types.MessageCreatedEvent, r.handleMessageCreated)
}
