package auth

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
)

type SignInCommandHandler interface {
	Handle(req *dto.SignInReqDTO) (*jwt.JWTResponse, *errorhandler.Response)
}

type signInCommandHandler struct {
	ctx                context.AppContext
	authDomainServices services.AuthDomainServices
}

func NewSignInCommandHandler(ctx context.AppContext, authDomainService services.AuthDomainServices) SignInCommandHandler {
	return &signInCommandHandler{
		ctx:                ctx,
		authDomainServices: authDomainService,
	}
}

func (s *signInCommandHandler) Handle(req *dto.SignInReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
	// SignIn the user
	res, err := s.authDomainServices.SignIn(req)

	if err != nil {
		return nil, err
	}

	return s.ctx.GetJwtContext().CreatePair(map[string]interface{}{
		"ID": res.ID,
	})
}
