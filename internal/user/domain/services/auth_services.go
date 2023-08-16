package services

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/entity"
	"github.com/iammuho/natternet/internal/user/domain/repository"
	"github.com/iammuho/natternet/internal/user/infrastructure/mongodb"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

type AuthDomainServices interface {
	Signin(req *dto.SigninReqDTO) (*entity.User, *errorhandler.Response)
}

type authDomainServices struct {
	ctx            context.AppContext
	userRepository repository.UserRepository
}

func NewAuthDomainServices(ctx context.AppContext) AuthDomainServices {
	// Initialize the repository
	userRepository := mongodb.NewUserRepository(ctx)

	return &authDomainServices{
		ctx:            ctx,
		userRepository: userRepository,
	}
}

func (a *authDomainServices) Signin(req *dto.SigninReqDTO) (*entity.User, *errorhandler.Response) {
	// Find the user by login (email or username)
	user, err := a.userRepository.FindOneByLogin(req.Login)

	if err != nil {
		return nil, err
	}

	// Check the user password
	if user == nil || !user.ComparePassword(req.Password) {
		return nil, &errorhandler.Response{Code: errorhandler.InvalidCredentialsErrorCode, Message: errorhandler.InvalidCredentialsMessage, StatusCode: fiber.StatusUnauthorized}
	}

	return user, nil
}
