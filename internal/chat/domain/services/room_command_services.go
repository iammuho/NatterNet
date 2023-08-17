package services

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/entity"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/internal/chat/infrastructure/mongodb"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomCommandDomainServices interface {
	CreateRoom(*dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
}

type roomCommandDomainServices struct {
	ctx            context.AppContext
	roomRepository repository.RoomRepository
}

func NewRoomCommandDomainServices(ctx context.AppContext) RoomCommandDomainServices {
	// Initialize the repository
	roomRepository := mongodb.NewRoomRepository(ctx)

	return &roomCommandDomainServices{
		ctx:            ctx,
		roomRepository: roomRepository,
	}
}

// CreateRoom creates a new room
func (r *roomCommandDomainServices) CreateRoom(req *dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	roomType := entity.RoomTypePrivate
	// Prepare the room type
	if req.IsGroup {
		roomType = entity.RoomTypeGroup
	}
	// Create the room
	roomEntity := entity.NewRoom(*entity.NewRoomMeta(req.Name, req.Description), roomType)

	// Add users
	for _, user := range req.UserIDs {
		roomEntity.AddRoomUser(*entity.NewRoomUser(user, entity.RoomUserRoleAdmin))
	}

	// Create the room
	if err := r.roomRepository.Create(values.NewRoomDBValue(roomEntity)); err != nil {
		return nil, err
	}

	return values.NewRoomValueFromRoom(roomEntity), nil
}
