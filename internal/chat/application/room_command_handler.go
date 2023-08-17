package application

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomCommandHandler struct {
	ctx context.AppContext
}

func NewRoomCommandHandler(ctx context.AppContext) *RoomCommandHandler {
	return &RoomCommandHandler{
		ctx: ctx,
	}
}

func (s *RoomCommandHandler) CreateRoom(req *dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	// Initialize the authentication domain service
	roomCommandDomainServices := services.NewRoomCommandDomainServices(s.ctx)

	return roomCommandDomainServices.CreateRoom(req)
}
