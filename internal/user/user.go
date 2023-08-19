package user

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth"
	"github.com/iammuho/natternet/internal/user/application/user"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/internal/user/infrastructure/mongodb"
)

// Application represents the application context for user-related functionality.
type Application struct {
	AppContext           context.AppContext
	SigninCommandHandler auth.SignInCommandHandler
	SignupCommandHandler auth.SignUpCommandHandler
	UserQueryHandler     user.UserQueryHandler
}

// NewUserApplication initializes a new user application context with the given app context.
func NewUserApplication(ctx context.AppContext) *Application {
	// Setup the user repository
	userRepository := mongodb.NewUserRepository(ctx)

	// Setup the domain services
	authDomainServices := services.NewAuthDomainServices(ctx, userRepository)

	// Initialize the userQueryDomainServices
	userQueryDomainServices := services.NewUserQueryDomainServices(ctx, userRepository)

	// Setup the command handlers
	signinCommandHandler := auth.NewSignInCommandHandler(ctx, authDomainServices)
	signupCommandHandler := auth.NewSignUpCommandHandler(ctx, authDomainServices)
	userQueryHandler := user.NewUserQueryHandler(ctx, userQueryDomainServices)

	return &Application{
		AppContext:           ctx,
		SigninCommandHandler: signinCommandHandler,
		SignupCommandHandler: signupCommandHandler,
		UserQueryHandler:     userQueryHandler,
	}
}
