package mongodb

//go:generate mockgen -destination=mocks/mock_mongodb_contexter.go -package=mongodb_mock -source=interface.go

import "go.mongodb.org/mongo-driver/mongo"

type MongoDBContext interface {
	// GetConn returns the mongodb connection
	GetConn() *mongo.Client

	// GetDatabase returns the mongodb database
	GetDatabase() *mongo.Database

	// Close closes the mongodb connection
	Close() error

	// Ping pings the mongodb connection
	Ping() error
}
