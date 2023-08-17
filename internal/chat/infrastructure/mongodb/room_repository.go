package mongodb

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type roomRepository struct {
	ctx context.AppContext
}

func NewRoomRepository(ctx context.AppContext) repository.RoomRepository {
	return &roomRepository{
		ctx: ctx,
	}
}

// Create creates a room
func (u *roomRepository) Create(room *values.RoomDBValue) *errorhandler.Response {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("rooms")

	_, err := collection.InsertOne(u.ctx.GetContext(), room)

	if err != nil {
		return &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return nil
}
