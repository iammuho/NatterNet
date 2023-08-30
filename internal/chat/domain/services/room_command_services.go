package services

//go:generate mockgen -destination=mocks/mock_room_command_services.go -package=mockchatdomainservices -source=room_command_services.go

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/entity"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	websocketTypes "github.com/iammuho/natternet/internal/user/domain/event/types"
	websocketValues "github.com/iammuho/natternet/internal/user/domain/values/websocket"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomCommandDomainServices interface {
	CreateRoom(*dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
	UpdateLastMessage(string, *values.MessageValue) *errorhandler.Response
	JoinRoom(*dto.JoinRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
	LeaveRoom(*dto.LeaveRoomReqDTO) (*values.RoomValue, *errorhandler.Response)
}

type roomCommandDomainServices struct {
	ctx            context.AppContext
	roomRepository repository.RoomRepository
}

func NewRoomCommandDomainServices(ctx context.AppContext, roomRepository repository.RoomRepository) RoomCommandDomainServices {
	return &roomCommandDomainServices{
		ctx:            ctx,
		roomRepository: roomRepository,
	}
}

// CreateRoom creates a new room
func (r *roomCommandDomainServices) CreateRoom(req *dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	// Create a user entity
	uuid := r.ctx.GetUUID().NewUUID()
	createdAt := r.ctx.GetTimer().Now()

	roomType := entity.RoomTypePrivate
	// Prepare the room type
	if req.IsGroup {
		roomType = entity.RoomTypeGroup
	}
	// Create the room
	roomEntity := entity.NewRoom(uuid, *entity.NewRoomMeta(req.Name, req.Description), roomType, createdAt)

	// Add the owner to the room as an admin
	roomEntity.AddRoomUser(*entity.NewRoomUser(req.Owner, entity.RoomUserRoleAdmin, createdAt))

	// Add other users as member
	for _, user := range req.UserIDs {
		roomEntity.AddRoomUser(*entity.NewRoomUser(user, entity.RoomUserRoleMember, createdAt))
	}

	// Create the room
	if err := r.roomRepository.Create(values.NewRoomDBValue(roomEntity)); err != nil {
		return nil, err
	}

	return values.NewRoomValueFromRoom(roomEntity), nil
}

// UpdateLastMessage updates the last message of the room
func (r *roomCommandDomainServices) UpdateLastMessage(roomID string, message *values.MessageValue) *errorhandler.Response {
	// Query the room
	room, err := r.roomRepository.GetRoomByID(roomID)

	if err != nil {
		return err
	}

	if room == nil {
		return &errorhandler.Response{Code: errorhandler.RoomNotFoundErrorCode, Message: errorhandler.RoomNotFoundMessage, StatusCode: fiber.StatusNotFound}
	}

	// Convert the room to entity
	roomEntity := room.ToRoom()

	// Update the last message
	now := r.ctx.GetTimer().Now()
	roomEntity.SetLastMessage(message.ID)
	roomEntity.SetUpdatedAt(&now)

	// Update the room
	return r.roomRepository.Update(values.NewRoomDBValue(roomEntity))
}

// JoinRoom joins a user to a room
func (r *roomCommandDomainServices) JoinRoom(req *dto.JoinRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	// Query the room
	room, err := r.roomRepository.GetRoomByID(req.RoomID)

	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, &errorhandler.Response{Code: errorhandler.RoomNotFoundErrorCode, Message: errorhandler.RoomNotFoundMessage, StatusCode: fiber.StatusNotFound}
	}

	// Convert the room to entity
	roomEntity := room.ToRoom()

	// Check if the room type is group
	if roomEntity.GetRoomConfig().RoomType != entity.RoomTypeGroup {
		return nil, &errorhandler.Response{Code: errorhandler.RoomNotGroupErrorCode, Message: errorhandler.RoomNotGroupMessage, StatusCode: fiber.StatusBadRequest}
	}

	// Check if the user is already in the room
	if roomEntity.CheckRoomUserExists(req.UserID) {
		return nil, &errorhandler.Response{Code: errorhandler.UserAlreadyInRoomErrorCode, Message: errorhandler.UserAlreadyInRoomMessage, StatusCode: fiber.StatusBadRequest}
	}

	// Add the user to the room
	roomEntity.AddRoomUser(*entity.NewRoomUser(req.UserID, entity.RoomUserRoleMember, r.ctx.GetTimer().Now()))

	// Update the room
	if err := r.roomRepository.Update(values.NewRoomDBValue(roomEntity)); err != nil {
		return nil, err
	}

	// Publishes to user websocket
	websocketEventValue := websocketValues.RoomUserJoinedWebsocketValue{
		UserID: req.UserID,
		RoomID: roomEntity.GetID(),
	}
	// Loop users and add them to the event with userid
	for _, user := range roomEntity.GetRoomUsers() {
		websocketEventValue.Users = append(websocketEventValue.Users, user.UserID)
	}

	messageJSON, _ := json.Marshal(websocketEventValue)

	_, publishErr := r.ctx.GetNatsContext().GetJetStreamContext().Publish(websocketTypes.RoomUserJoinedEvent, messageJSON)

	if publishErr != nil {
		r.ctx.GetLogger().Error(publishErr.Error())
	}

	return values.NewRoomValueFromRoom(roomEntity), nil
}

// LeaveRoom removes a user from a room
func (r *roomCommandDomainServices) LeaveRoom(req *dto.LeaveRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	// Query the room
	room, err := r.roomRepository.GetRoomByID(req.RoomID)

	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, &errorhandler.Response{Code: errorhandler.RoomNotFoundErrorCode, Message: errorhandler.RoomNotFoundMessage, StatusCode: fiber.StatusNotFound}
	}

	// Convert the room to entity
	roomEntity := room.ToRoom()

	// Check if the user is in the room
	if !roomEntity.CheckRoomUserExists(req.UserID) {
		return nil, &errorhandler.Response{Code: errorhandler.UserIsNotInRoomCode, Message: errorhandler.UserIsNotInRoomMessage, StatusCode: fiber.StatusBadRequest}
	}

	// Remove the user from the room
	roomEntity.RemoveRoomUser(req.UserID)

	// Update the room
	if err := r.roomRepository.Update(values.NewRoomDBValue(roomEntity)); err != nil {
		return nil, err
	}

	return values.NewRoomValueFromRoom(roomEntity), nil
}
