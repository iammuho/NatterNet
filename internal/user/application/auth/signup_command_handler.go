package auth

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
)

type SignUpCommandHandler struct {
	ctx context.AppContext
}

func NewSignUpCommandHandler(ctx context.AppContext) *SignUpCommandHandler {
	return &SignUpCommandHandler{
		ctx: ctx,
	}
}

func (s *SignUpCommandHandler) Handle(req *dto.SignupReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
	// Initialize the authentication domain service
	authDomainService := services.NewAuthDomainServices(s.ctx)

	// SignUp the user
	res, err := authDomainService.SignUp(req)

	if err != nil {
		return nil, err
	}

	return s.ctx.GetJwtContext().CreatePair(map[string]interface{}{
		"ID": res.ID,
	})
}
