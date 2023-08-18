package application

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomCommandHandler struct {
	ctx                 context.AppContext
	roomCommandServices services.RoomCommandDomainServices
}

func NewRoomCommandHandler(ctx context.AppContext, roomCommandServices services.RoomCommandDomainServices) *RoomCommandHandler {
	return &RoomCommandHandler{
		ctx:                 ctx,
		roomCommandServices: roomCommandServices,
	}
}

func (r *RoomCommandHandler) CreateRoom(req *dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	return r.roomCommandServices.CreateRoom(req)
}
