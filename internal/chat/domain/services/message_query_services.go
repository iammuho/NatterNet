package services

//go:generate mockgen -destination=mocks/mock_message_query_services.go -package=mockchatdomainservices -source=message_query_services.go

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

type MessageQueryDomainServices interface {
	QueryMessages(*dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response)
}

type messageQueryDomainServices struct {
	ctx                     context.AppContext
	messageRepository       repository.MessageRepository
	roomQueryDomainServices RoomQueryDomainServices
}

func NewMessageQueryDomainServices(ctx context.AppContext, messageRepository repository.MessageRepository, roomQueryDomainServices RoomQueryDomainServices) MessageQueryDomainServices {
	return &messageQueryDomainServices{
		ctx:                     ctx,
		messageRepository:       messageRepository,
		roomQueryDomainServices: roomQueryDomainServices,
	}
}

// CreateMessage creates a new message
func (r *messageQueryDomainServices) QueryMessages(req *dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response) {
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
	if !roomEntity.CheckRoomUserExists(req.UserID) {
		return nil, &errorhandler.Response{Code: errorhandler.UserIsNotInRoomCode, Message: errorhandler.UserIsNotInRoomMessage, StatusCode: fiber.StatusForbidden}
	}

	// Query the messages
	messages, err := r.messageRepository.Query(req)

	if err != nil {
		return nil, err
	}

	return messages, nil
}
