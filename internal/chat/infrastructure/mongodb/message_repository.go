package mongodb

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/repository"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Query queries messages
func (u *messageRepository) Query(req *dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response) {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("messages")

	// Filters
	filters := bson.D{
		{Key: "room_id", Value: req.RoomID},
	}

	// Sorting
	sort := bson.D{}

	sortField := "created_at"
	sortOrder := -1

	if req.SortField != "" {
		sortField = req.SortField
		if req.SortOrder != "" {
			switch req.SortOrder {
			case "asc":
				sortOrder = 1
			case "desc":
				sortOrder = -1
			}
		}
	}

	sort = append(sort, bson.E{Key: sortField, Value: sortOrder})

	// Pagination
	skip := int64((req.Page - 1) * req.PerPage)
	limit := int64(req.PerPage)

	// Query
	cursor, err := collection.Find(u.ctx.GetContext(), filters, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  sort,
	})

	if err != nil {
		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	// Decode
	var messages []*values.MessageValue
	if err := cursor.All(u.ctx.GetContext(), &messages); err != nil {
		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return messages, nil
}
