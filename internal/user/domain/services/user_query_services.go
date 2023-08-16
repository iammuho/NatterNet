package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/domain/repository"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/internal/user/infrastructure/mongodb"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type UserQueryDomainServices interface {
	FindByID(id string) (*values.UserValue, *errorhandler.Response)
}

type userQueryDomainServices struct {
	ctx            context.AppContext
	userRepository repository.UserRepository
}

func NewUserQueryDomainServices(ctx context.AppContext) UserQueryDomainServices {
	// Initialize the repository
	userRepository := mongodb.NewUserRepository(ctx)

	return &userQueryDomainServices{
		ctx:            ctx,
		userRepository: userRepository,
	}
}

func (s *userQueryDomainServices) FindByID(id string) (*values.UserValue, *errorhandler.Response) {
	user, err := s.userRepository.FindOneByID(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errorhandler.Response{Code: errorhandler.UserNotFoundErrorCode, Message: errorhandler.UserNotFoundMessage, StatusCode: fiber.StatusNotFound}
	}

	return user.ToUserValue(), nil

}
