package repository

//go:generate mockgen -destination=mocks/mock_room_repository.go -package=mockchatrepository -source=room_repository.go

import (
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomRepository interface {
	// Commands
	Create(room *values.RoomDBValue) *errorhandler.Response

	// Queries
	QueryRooms(req *dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response)
	GetRoomByID(id string) (*values.RoomValue, *errorhandler.Response)
}
