package application

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type MessageCommandHandler interface {
	CreateMessage(*dto.CreateMessageReqDTO) (*values.MessageValue, *errorhandler.Response)
}

type messageCommandHandler struct {
	ctx                    context.AppContext
	messageCommandServices services.MessageCommandDomainServices
}

func NewMessageCommandHandler(ctx context.AppContext, messageCommandServices services.MessageCommandDomainServices) MessageCommandHandler {
	return &messageCommandHandler{
		ctx:                    ctx,
		messageCommandServices: messageCommandServices,
	}
}

func (r *messageCommandHandler) CreateMessage(req *dto.CreateMessageReqDTO) (*values.MessageValue, *errorhandler.Response) {
	return r.messageCommandServices.CreateMessage(req)
}
