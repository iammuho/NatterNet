package application

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/services"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type MessageQueryHandler interface {
	QueryMessages(*dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response)
}

type messageQueryHandler struct {
	ctx                  context.AppContext
	messageQueryServices services.MessageQueryDomainServices
}

func NewMessageQueryHandler(ctx context.AppContext, messageQueryServices services.MessageQueryDomainServices) MessageQueryHandler {
	return &messageQueryHandler{
		ctx:                  ctx,
		messageQueryServices: messageQueryServices,
	}
}

func (r *messageQueryHandler) QueryMessages(req *dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response) {
	return r.messageQueryServices.QueryMessages(req)
}
