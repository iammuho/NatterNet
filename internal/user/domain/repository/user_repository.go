package repository

import (
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type UserRepository interface {
	FindOneByLogin(login string) (*values.UserDBValue, *errorhandler.Response)
	FindOneByEmail(email string) (*values.UserDBValue, *errorhandler.Response)
	FindOneByUsername(username string) (*values.UserDBValue, *errorhandler.Response)

	// Commands
	Create(user *values.UserDBValue) *errorhandler.Response
}
