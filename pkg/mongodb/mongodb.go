// Package mongodb is a module for mongodb
package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type server struct {
	conn     *mongo.Client
	database *mongo.Database
}

// NewMongoDB creates a new mongodb connection
func NewMongoDB(opts ...Option) (MongoDBContext, error) {
	mongodbOptions := MongoDBOptions{}
	for _, o := range opts {
		o(&mongodbOptions)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbOptions.URI).SetAuth(options.Credential{
		Username: mongodbOptions.Username,
		Password: mongodbOptions.Password,
	}))
	if err != nil {
		return nil, err
	}

	// ping
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &server{
		conn:     client,
		database: client.Database(mongodbOptions.Database),
	}, nil
}

// GetConn returns the mongodb connection
func (s *server) GetConn() *mongo.Client {
	return s.conn
}

// GetDatabase returns the mongodb database
func (s *server) GetDatabase() *mongo.Database {
	return s.database
}

// Close closes the mongodb connection
func (s *server) Close() error {
	return s.conn.Disconnect(context.Background())
}

// Ping pings the mongodb connection
func (s *server) Ping() error {
	return s.conn.Ping(context.Background(), nil)
}
