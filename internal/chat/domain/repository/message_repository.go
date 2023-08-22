package repository

//go:generate mockgen -destination=mocks/mock_message_repository.go -package=mockchatrepository -source=message_repository.go

import (
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type MessageRepository interface {
	// Commands
	Create(message *values.MessageDBValue) *errorhandler.Response

	// Queries
	Query(query *dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response)
}
