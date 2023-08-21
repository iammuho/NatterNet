package services

//go:generate mockgen -destination=mocks/mock_room_query_services.go -package=mockchatdomainservices -source=room_query_services.go

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomQueryDomainServices interface {
	QueryRooms(*dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response)
	GetRoomByID(string) (*values.RoomValue, *errorhandler.Response)
}

type roomQueryDomainServices struct {
	ctx            context.AppContext
	roomRepository repository.RoomRepository
}

func NewRoomQueryDomainServices(ctx context.AppContext, roomRepository repository.RoomRepository) RoomQueryDomainServices {
	return &roomQueryDomainServices{
		ctx:            ctx,
		roomRepository: roomRepository,
	}
}

// CreateRoom creates a new room
func (r *roomQueryDomainServices) QueryRooms(req *dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response) {
	return r.roomRepository.QueryRooms(req)
}

// GetRoomByID gets a room by id
func (r *roomQueryDomainServices) GetRoomByID(id string) (*values.RoomValue, *errorhandler.Response) {
	return r.roomRepository.GetRoomByID(id)
}
