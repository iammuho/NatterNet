package context

import (
	contextPKG "context"

	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
)

type AppContext interface {
	GetContext() contextPKG.Context
	GetLogger() *logger.Logger
	GetJwtContext() jwt.JwtContext
	GetMongoContext() mongodb.MongoDBContext
}
