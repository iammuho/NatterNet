package auth

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
)

type SigninQueryHandler struct {
	ctx context.AppContext
}

func NewSigninQueryHandler(ctx context.AppContext) *SigninQueryHandler {
	return &SigninQueryHandler{
		ctx: ctx,
	}
}

func (s *SigninQueryHandler) Handle(req *dto.SigninReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
	return s.ctx.GetJwtContext().CreatePair(map[string]interface{}{
		"ID": "1",
	})
}
