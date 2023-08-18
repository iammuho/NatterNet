package services

//go:generate mockgen -destination=mocks/mock_auth_domain_services.go -package=mockuserdomainservices -source=auth_services.go

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/entity"
	"github.com/iammuho/natternet/internal/user/domain/repository"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
)

type AuthDomainServices interface {
	SignIn(req *dto.SignInReqDTO) (*values.UserValue, *errorhandler.Response)
	SignUp(req *dto.SignupReqDTO) (*values.UserValue, *errorhandler.Response)
}

type authDomainServices struct {
	ctx            context.AppContext
	userRepository repository.UserRepository
}

func NewAuthDomainServices(ctx context.AppContext, userRepository repository.UserRepository) AuthDomainServices {
	return &authDomainServices{
		ctx:            ctx,
		userRepository: userRepository,
	}
}

// Signin is the domain service for the signin route
func (a *authDomainServices) SignIn(req *dto.SignInReqDTO) (*values.UserValue, *errorhandler.Response) {
	// Find the user by login (email or username)
	user, err := a.userRepository.FindOneByLogin(req.Login)

	if err != nil {
		return nil, err
	}

	// Convert the db to entity
	userEntity := user.ToUserEntity()

	// Check the user password
	if user == nil || !a.ctx.GetHashingFactory().ComparePassword(userEntity.GetPassword(), req.Password) {
		return nil, &errorhandler.Response{Code: errorhandler.InvalidCredentialsErrorCode, Message: errorhandler.InvalidCredentialsMessage, StatusCode: fiber.StatusUnauthorized}
	}

	return values.NewUserValueFromUser(userEntity), nil
}

// SignUp is the domain service for the signup route
func (a *authDomainServices) SignUp(req *dto.SignupReqDTO) (*values.UserValue, *errorhandler.Response) {
	// Check if the email already exists
	user, err := a.userRepository.FindOneByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, &errorhandler.Response{Code: errorhandler.EmailAlreadyExistsErrorCode, Message: errorhandler.EmailAlreadyExistsMessage, StatusCode: fiber.StatusConflict}
	}

	// Check if the username already exists
	user, err = a.userRepository.FindOneByUsername(req.Username)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, &errorhandler.Response{Code: errorhandler.UsernameAlreadyExistsErrorCode, Message: errorhandler.UsernameAlreadyExistsMessage, StatusCode: fiber.StatusConflict}
	}

	// Create a user entity
	uuid := a.ctx.GetUUID().NewUUID()
	createdAt := a.ctx.GetTimer().Now()

	userEntity := entity.NewUser(uuid, req.Username, req.Password, req.Email, createdAt)

	// Create the user
	err = a.userRepository.Create(values.NewUserDBValueFromUser(userEntity))

	if err != nil {
		return nil, err
	}

	return values.NewUserValueFromUser(userEntity), nil
}
