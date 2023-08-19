package auth

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
)

type SignUpCommandHandler interface {
	Handle(req *dto.SignupReqDTO) (*jwt.JWTResponse, *errorhandler.Response)
}

type signUpCommandHandler struct {
	ctx               context.AppContext
	authDomainService services.AuthDomainServices
}

func NewSignUpCommandHandler(ctx context.AppContext, authDomainService services.AuthDomainServices) SignUpCommandHandler {
	return &signUpCommandHandler{
		ctx:               ctx,
		authDomainService: authDomainService,
	}
}

func (s *signUpCommandHandler) Handle(req *dto.SignupReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
	res, err := s.authDomainService.SignUp(req)

	if err != nil {
		return nil, err
	}

	return s.ctx.GetJwtContext().CreatePair(map[string]interface{}{
		"ID": res.ID,
	})
}
