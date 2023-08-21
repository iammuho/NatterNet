package application

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomQueryHandler interface {
	QueryRooms(*dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response)
}

type roomQueryHandler struct {
	ctx               context.AppContext
	roomQueryServices services.RoomQueryDomainServices
}

func NewRoomQueryHandler(ctx context.AppContext, roomQueryServices services.RoomQueryDomainServices) RoomQueryHandler {
	return &roomQueryHandler{
		ctx:               ctx,
		roomQueryServices: roomQueryServices,
	}
}

// Query rooms
func (r *roomQueryHandler) QueryRooms(req *dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response) {
	return r.roomQueryServices.QueryRooms(req)
}
