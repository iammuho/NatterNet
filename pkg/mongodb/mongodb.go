// Package mongodb is a module for mongodb
package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Server is a structure which Nats server client
type server struct {
	Conn     *mongo.Client
	Database *mongo.Database
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
		Conn:     client,
		Database: client.Database(mongodbOptions.Database),
	}, nil
}

// GetConn returns the mongodb connection
func (s *server) GetConn() *mongo.Client {
	return s.Conn
}

// GetDatabase returns the mongodb database
func (s *server) GetDatabase() *mongo.Database {
	return s.Database
}

// Close closes the mongodb connection
func (s *server) Close() error {
	return s.Conn.Disconnect(context.Background())
}

// Ping pings the mongodb connection
func (s *server) Ping() error {
	return s.Conn.Ping(context.Background(), nil)
}
