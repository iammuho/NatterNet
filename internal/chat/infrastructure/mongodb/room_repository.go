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

// QueryRooms queries rooms
func (u *roomRepository) QueryRooms(req *dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response) {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("rooms")

	// Filters
	filters := bson.D{}

	if len(req.UserIn) > 0 {
		filters = append(filters, bson.E{Key: "users.user_id", Value: bson.D{{Key: "$in", Value: req.UserIn}}})
	}

	if len(req.UserNotIn) > 0 {
		filters = append(filters, bson.E{Key: "users.user_id", Value: bson.D{{Key: "$nin", Value: req.UserNotIn}}})
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
	var rooms []*values.RoomValue

	if err := cursor.All(u.ctx.GetContext(), &rooms); err != nil {
		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return rooms, nil
}
