package services

//go:generate mockgen -destination=mocks/mock_message_command_services.go -package=mockchatdomainservices -source=message_command_services.go

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/entity"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

type MessageCommandDomainServices interface {
	CreateMessage(*dto.CreateMessageReqDTO) (*values.MessageValue, *errorhandler.Response)
}

type messageCommandDomainServices struct {
	ctx                     context.AppContext
	messageRepository       repository.MessageRepository
	roomQueryDomainServices RoomQueryDomainServices
}

func NewMessageCommandDomainServices(ctx context.AppContext, messageRepository repository.MessageRepository, roomQueryDomainServices RoomQueryDomainServices) MessageCommandDomainServices {
	return &messageCommandDomainServices{
		ctx:                     ctx,
		messageRepository:       messageRepository,
		roomQueryDomainServices: roomQueryDomainServices,
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

	// Create a user entity
	uuid := r.ctx.GetUUID().NewUUID()
	createdAt := r.ctx.GetTimer().Now()

	// Create the message
	messageEntity := entity.NewMessage(uuid, req.RoomID, req.SenderID, req.Content, entity.MessageType(req.MessageType), createdAt)

	// Create the message
	if err := r.messageRepository.Create(values.NewMessageDBValueFromMessage(messageEntity)); err != nil {
		return nil, err
	}

	return values.NewMessageValueFromMessage(messageEntity), nil
}
