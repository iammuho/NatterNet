package user

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/user/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type UserQueryHandler struct {
	ctx context.AppContext
}

func NewUserQueryHandler(ctx context.AppContext) *UserQueryHandler {
	return &UserQueryHandler{
		ctx: ctx,
	}
}

func (s *UserQueryHandler) QueryUserByID(req *dto.QueryUserByIDReqDTO) (*values.UserValue, *errorhandler.Response) {
	// Initialize the user query domain service
	userQueryDomainServices := services.NewUserQueryDomainServices(s.ctx)

	// Find the user
	return userQueryDomainServices.FindByID(req.UserID)
}
