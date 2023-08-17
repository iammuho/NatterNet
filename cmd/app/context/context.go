package context

import (
	contextPKG "context"

	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
)

//go:generate mockgen -destination=mocks/mock_app_contexter.go -package=mockcontext -source=context.go

type AppContext interface {
	GetContext() contextPKG.Context
	GetLogger() *logger.Logger
	GetJwtContext() jwt.JwtContext
	GetMongoContext() mongodb.MongoDBContext
}
