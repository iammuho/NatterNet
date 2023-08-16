package repository

import (
	"github.com/iammuho/natternet/internal/user/domain/entity"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type UserRepository interface {
	FindOneByLogin(login string) (*entity.User, *errorhandler.Response)
}
