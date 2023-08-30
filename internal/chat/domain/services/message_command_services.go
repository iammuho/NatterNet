package services

//go:generate mockgen -destination=mocks/mock_message_command_services.go -package=mockchatdomainservices -source=message_command_services.go

import (
	"encoding/json"

	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/entity"
	"github.com/iammuho/natternet/internal/chat/domain/event/types"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	websocketTypes "github.com/iammuho/natternet/internal/user/domain/event/types"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

type MessageCommandDomainServices interface {
	CreateMessage(*dto.CreateMessageReqDTO) (*values.MessageValue, *errorhandler.Response)
}

type messageCommandDomainServices struct {
	ctx                       context.AppContext
	messageRepository         repository.MessageRepository
	roomQueryDomainServices   RoomQueryDomainServices
	roomCommandDomainServices RoomCommandDomainServices
}

func NewMessageCommandDomainServices(ctx context.AppContext, messageRepository repository.MessageRepository, roomQueryDomainServices RoomQueryDomainServices, roomCoomroomCommandDomainServices RoomCommandDomainServices) MessageCommandDomainServices {
	return &messageCommandDomainServices{
		ctx:                       ctx,
		messageRepository:         messageRepository,
		roomQueryDomainServices:   roomQueryDomainServices,
		roomCommandDomainServices: roomCoomroomCommandDomainServices,
	}
}

// CreateMessage creates a new message
func (r *messageCommandDomainServices) CreateMessage(req *dto.CreateMessageReqDTO) (*values.MessageValue, *errorhandler.Response) {
	// Query the room to make sure it exists
	room, err := r.roomQueryDomainServices.GetRoomByID(req.RoomID)

	if err != nil {
		return nil, err
	}

	if room == nil {
		return nil, &errorhandler.Response{Code: errorhandler.RoomNotFoundErrorCode, Message: errorhandler.RoomNotFoundMessage, StatusCode: fiber.StatusNotFound}
	}

	// Convert the room to entity
	roomEntity := room.ToRoom()

	// Make sure the sender is in the room
	if !roomEntity.CheckRoomUserExists(req.SenderID) {
		return nil, &errorhandler.Response{Code: errorhandler.UserIsNotInRoomCode, Message: errorhandler.UserIsNotInRoomMessage, StatusCode: fiber.StatusForbidden}
	}

	// Check if message type is valid
	if !entity.IsValidMessageType(req.MessageType) {
		return nil, &errorhandler.Response{Code: errorhandler.InvalidMessageTypeCode, Message: errorhandler.InvalidMessageTypeMessage, StatusCode: fiber.StatusBadRequest}
	}

	// Create a user entity
	uuid := r.ctx.GetUUID().NewUUID()
	createdAt := r.ctx.GetTimer().Now()

	// Create the message
	messageEntity := entity.NewMessage(uuid, req.RoomID, req.SenderID, req.Content, entity.MessageType(req.MessageType), createdAt)

	// Create the message
	if err := r.messageRepository.Create(values.NewMessageDBValueFromMessage(messageEntity)); err != nil {
		return nil, err
	}

	// publish to nats
	messageJSON, _ := json.Marshal(values.NewMessageValueFromMessage(messageEntity))
	_, publishErr := r.ctx.GetNatsContext().GetJetStreamContext().Publish(types.MessageCreatedEvent, messageJSON)

	if publishErr != nil {
		r.ctx.GetLogger().Error(publishErr.Error())
	}

	// Publishes to user websocket
	_, publishErr = r.ctx.GetNatsContext().GetJetStreamContext().Publish(websocketTypes.MessageCreatedEvent, messageJSON)

	if publishErr != nil {
		r.ctx.GetLogger().Error(publishErr.Error())
	}

	return values.NewMessageValueFromMessage(messageEntity), nil
}
