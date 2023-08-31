package application

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomCommandHandler interface {
	CreateRoom(*dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
	JoinRoom(*dto.JoinRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
	LeaveRoom(*dto.LeaveRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
	SendRoomEvent(*dto.SendRoomEventReqDTO) *errorhandler.Response
}

type roomCommandHandler struct {
	ctx                 context.AppContext
	roomCommandServices services.RoomCommandDomainServices
}

func NewRoomCommandHandler(ctx context.AppContext, roomCommandServices services.RoomCommandDomainServices) RoomCommandHandler {
	return &roomCommandHandler{
		ctx:                 ctx,
		roomCommandServices: roomCommandServices,
	}
}

func (r *roomCommandHandler) CreateRoom(req *dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	return r.roomCommandServices.CreateRoom(req)
}

func (r *roomCommandHandler) JoinRoom(req *dto.JoinRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	return r.roomCommandServices.JoinRoom(req)
}

func (r *roomCommandHandler) LeaveRoom(req *dto.LeaveRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	return r.roomCommandServices.LeaveRoom(req)
}

func (r *roomCommandHandler) SendRoomEvent(req *dto.SendRoomEventReqDTO) *errorhandler.Response {
	return r.roomCommandServices.SendRoomEvent(req)
}
