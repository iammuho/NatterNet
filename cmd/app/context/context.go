package context

import (
	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
)

type AppContext interface {
	GetContext() interface{}
	GetLogger() *logger.Logger
	GetJwtContext() jwt.JwtContext
	GetMongoContext() mongodb.MongoDBContext
}
