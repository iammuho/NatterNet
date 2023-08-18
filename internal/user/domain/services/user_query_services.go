package services

//go:generate mockgen -destination=mocks/mock_user_query_domain_services.go -package=mockuserdomainservices -source=user_query_services.go

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/domain/repository"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type UserQueryDomainServices interface {
	FindByID(id string) (*values.UserValue, *errorhandler.Response)
}

type userQueryDomainServices struct {
	ctx            context.AppContext
	userRepository repository.UserRepository
}

func NewUserQueryDomainServices(ctx context.AppContext, userRepository repository.UserRepository) UserQueryDomainServices {
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
