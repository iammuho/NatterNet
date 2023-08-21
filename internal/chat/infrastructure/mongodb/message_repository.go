package mongodb

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type messageRepository struct {
	ctx context.AppContext
}

func NewMessageRepository(ctx context.AppContext) repository.MessageRepository {
	return &messageRepository{
		ctx: ctx,
	}
}

// Create creates a message
func (u *messageRepository) Create(message *values.MessageDBValue) *errorhandler.Response {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("messages")

	_, err := collection.InsertOne(u.ctx.GetContext(), message)

	if err != nil {
		return &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return nil
}
