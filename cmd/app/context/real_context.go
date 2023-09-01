package context

import (
	"context"

	"github.com/iammuho/natternet/pkg/hashing"
	"github.com/iammuho/natternet/pkg/jwt"
	"github.com/iammuho/natternet/pkg/logger"
	"github.com/iammuho/natternet/pkg/mongodb"
	"github.com/iammuho/natternet/pkg/nats"
	"github.com/iammuho/natternet/pkg/storage"
	"github.com/iammuho/natternet/pkg/utils"
)

type appContext struct {
	ctx            context.Context
	logger         *logger.Logger
	jwt            jwt.JwtContext
	mongoContext   mongodb.MongoDBContext
	hashingFactory hashing.HashingFactory
	natsContext    nats.NatsContext
	storageContext storage.StorageContext
	UUID           utils.UUID
	Timer          utils.Timer
}

func NewAppContext(logger *logger.Logger, jwt jwt.JwtContext, mongoContext mongodb.MongoDBContext, natsContext nats.NatsContext, storageContext storage.StorageContext) AppContext {
	ctx := context.Background()

	// Set the UUID
	uuid := utils.RealUUID{}

	// Set the timer
	timer := utils.RealTimer{}

	// Set the hashing factory
	hashingFactory := hashing.NewHashingFactory()

	return &appContext{
		ctx:            ctx,
		logger:         logger,
		jwt:            jwt,
		mongoContext:   mongoContext,
		hashingFactory: hashingFactory,
		natsContext:    natsContext,
		storageContext: storageContext,
		UUID:           uuid,
		Timer:          timer,
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

func (c *appContext) GetUUID() utils.UUID {
	return c.UUID
}

func (c *appContext) GetTimer() utils.Timer {
	return c.Timer
}

func (c *appContext) GetHashingFactory() hashing.HashingFactory {
	return c.hashingFactory
}

func (c *appContext) GetNatsContext() nats.NatsContext {
	return c.natsContext
}

func (c *appContext) GetStorageContext() storage.StorageContext {
	return c.storageContext
}
