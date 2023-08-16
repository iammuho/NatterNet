package context

import (
	"context"

	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
)

type appContext struct {
	ctx          context.Context
	logger       *logger.Logger
	mongoContext mongodb.MongoDBContext
}

func NewAppContext(logger *logger.Logger, mongoContext mongodb.MongoDBContext) AppContext {
	ctx := context.Background()

	return &appContext{
		ctx:          ctx,
		logger:       logger,
		mongoContext: mongoContext,
	}
}

func (c *appContext) GetContext() interface{} {
	return c.ctx
}

func (c *appContext) GetLogger() *logger.Logger {
	return c.logger
}

func (c *appContext) GetMongoContext() mongodb.MongoDBContext {
	return c.mongoContext
}
