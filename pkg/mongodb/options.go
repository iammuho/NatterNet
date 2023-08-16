package mongodb

// Option is the func interface to assign options
type Option func(*MongoDBOptions)

// MongoDBOptions defines the mongo configuration
type MongoDBOptions struct {
	URI      string
	Database string
	Username string
	Password string
}

// WithMongoDBURI defines the Mongo Connection URI
func WithMongoDBURI(uri string) Option {
	return func(o *MongoDBOptions) {
		o.URI = uri
	}
}

// WithMongoDBDatabase defines the Mongo Connection Database
func WithMongoDBDatabase(database string) Option {
	return func(o *MongoDBOptions) {
		o.Database = database
	}
}

// WithMongoDBUsername defines the Mongo Connection Username
func WithMongoDBUsername(username string) Option {
	return func(o *MongoDBOptions) {
		o.Username = username
	}
}

// WithMongoDBPassword defines the Mongo Connection password
func WithMongoDBPassword(password string) Option {
	return func(o *MongoDBOptions) {
		o.Password = password
	}
}
