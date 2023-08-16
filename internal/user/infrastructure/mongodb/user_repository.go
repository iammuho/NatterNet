package mongodb

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/domain/entity"
	"github.com/iammuho/natternet/internal/user/domain/repository"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	ctx context.AppContext
}

func NewUserRepository(ctx context.AppContext) repository.UserRepository {
	return &userRepository{
		ctx: ctx,
	}
}

func (u *userRepository) FindOneByLogin(login string) (*entity.User, *errorhandler.Response) {
	// Get the collection
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("users")

	// Find the user by login (email or username)
	var user entity.User
	err := collection.FindOne(u.ctx.GetContext(), bson.M{"$or": []bson.M{{"email": login}, {"username": login}}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return &user, nil
}
