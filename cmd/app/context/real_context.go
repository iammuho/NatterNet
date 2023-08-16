package context

import (
	"context"

	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
)

type appContext struct {
	ctx          context.Context
	logger       *logger.Logger
	jwt          jwt.JwtContext
	mongoContext mongodb.MongoDBContext
}

func NewAppContext(logger *logger.Logger, jwt jwt.JwtContext, mongoContext mongodb.MongoDBContext) AppContext {
	ctx := context.Background()

	return &appContext{
		ctx:          ctx,
		logger:       logger,
		jwt:          jwt,
		mongoContext: mongoContext,
	}
}

func (c *appContext) GetContext() context.Context {
	return c.ctx
}

func (c *appContext) GetLogger() *logger.Logger {
	return c.logger
}

func (c *appContext) GetJwtContext() jwt.JwtContext {
	return c.jwt
}

func (c *appContext) GetMongoContext() mongodb.MongoDBContext {
	return c.mongoContext
}
