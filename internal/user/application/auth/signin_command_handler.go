package auth

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
)

type SignInCommandHandler struct {
	ctx context.AppContext
}

func NewSignInCommandHandler(ctx context.AppContext) *SignInCommandHandler {
	return &SignInCommandHandler{
		ctx: ctx,
	}
}

func (s *SignInCommandHandler) Handle(req *dto.SignInReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
	// Initialize the authentication domain service
	authDomainService := services.NewAuthDomainServices(s.ctx)

	// SignIn the user
	res, err := authDomainService.SignIn(req)

	if err != nil {
		return nil, err
	}

	return s.ctx.GetJwtContext().CreatePair(map[string]interface{}{
		"ID": res.ID,
	})
}