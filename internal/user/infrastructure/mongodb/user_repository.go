package mongodb

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/domain/repository"
	"github.com/iammuho/natternet/internal/user/domain/values"
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

// FindOneByLogin finds a user by login (email or username)
func (u *userRepository) FindOneByLogin(login string) (*values.UserDBValue, *errorhandler.Response) {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("users")

	var user values.UserDBValue
	err := collection.FindOne(u.ctx.GetContext(), bson.M{"$or": []bson.M{{"email": login}, {"username": login}}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return &user, nil
}

// FindOneByEmail finds a user by email
func (u *userRepository) FindOneByEmail(email string) (*values.UserDBValue, *errorhandler.Response) {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("users")

	var user values.UserDBValue
	err := collection.FindOne(u.ctx.GetContext(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return &user, nil
}

// FindOneByUsername finds a user by username
func (u *userRepository) FindOneByUsername(username string) (*values.UserDBValue, *errorhandler.Response) {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("users")

	var user values.UserDBValue
	err := collection.FindOne(u.ctx.GetContext(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return &user, nil
}

// Create creates a user
func (u *userRepository) Create(user *values.UserDBValue) *errorhandler.Response {
	collection := u.ctx.GetMongoContext().GetDatabase().Collection("users")

	_, err := collection.InsertOne(u.ctx.GetContext(), user)

	if err != nil {
		return &errorhandler.Response{Code: errorhandler.DatabaseErrorCode, Message: err.Error()}
	}

	return nil
}
