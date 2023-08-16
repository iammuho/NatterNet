package auth

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
)

type SigninCommandHandler struct {
	ctx context.AppContext
}

func NewSigninCommandHandler(ctx context.AppContext) *SigninCommandHandler {
	return &SigninCommandHandler{
		ctx: ctx,
	}
}

func (s *SigninCommandHandler) Handle(req *dto.SigninReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
	// Initialize the authentication domain service
	authDomainService := services.NewAuthDomainServices(s.ctx)

	// Signin the user
	res, err := authDomainService.Signin(req)

	if err != nil {
		return nil, err
	}

	return s.ctx.GetJwtContext().CreatePair(map[string]interface{}{
		"ID": res.GetID(),
	})
}
