package config

import "time"

// Config stores configuration values
var Config *config

type config struct {

	// Application provides the application basic configurations.
	Application struct {
		Name        string `env:"APPLICATION_NAME"     envDefault:"NatterNet"`
		Environment string `env:"APPLICATION_ENV"     envDefault:"local"`
		Version     string `env:"APPLICATION_VERSION"     envDefault:"0.1"`
	}

	// Logger provides the logger basic configurations.
	Logger struct {
		Level string `env:"LOGGER_LEVEL"     envDefault:"debug"`
		Name  string `env:"LOGGER_NAME"     envDefault:"natternet"`
	}

	// HTTPServer provides the HTTP server configuration.
	HTTPServer struct {
		ListenAddress string `env:"SERVER_LISTEN_ADDRESS"     envDefault:"0.0.0.0"`
		ListenPort    int    `env:"SERVER_LISTEN_PORT"     envDefault:"8080"`

		CaseSensitive      bool          `env:"SERVER_CASE_SENSITIVE" envDefault:"true"`
		StrictRouting      bool          `env:"SERVER_STRICT_ROUTING" envDefault:"true"`
		ReadTimeout        time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"5s"`
		WriteTimeout       time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"5s"`
		MaxConnsPerIP      int           `env:"SERVER_MAX_CONN_PER_IP" envDefault:"50"`
		MaxRequestsPerConn int           `env:"SERVER_MAX_REQUESTS_PER_CONN" envDefault:"10"`
		BodyLimit          int           `env:"SERVER_BODY_LIMIT" envDefault:"4096"`

		TLSEnabled bool `env:"SERVER_TLS_ENABLED" envDefault:"false"`
	}

	// MongoDB provides the MongoDB configuration.
	MongoDB struct {
		URI      string `env:"MONGODB_URI" envDefault:"mongodb://localhost:27017"`
		Database string `env:"MONGODB_DATABASE" envDefault:"natternet"`
		Username string `env:"MONGODB_USERNAME" envDefault:"root"`
		Password string `env:"MONGODB_PASSWORD" envDefault:"natternet"`
	}

	// JWT provides the JWT configuration.
	JWT struct {
		PublicKeyPath  string `env:"JWT_PUBLIC_KEY_PATH" envDefault:"/etc/ssl/certs/client.pub.pem"`
		PrivateKeyPath string `env:"JWT_PRIVATE_KEY_PATH" envDefault:"/etc/ssl/certs/client.pem"`
		Kid            string `env:"JWT_KID" envDefault:"natternet"`
		Issuer         string `env:"JWT_ISSUER" envDefault:"natternet"`
		Subject        string `env:"JWT_SUBJECT" envDefault:"natternet"`
	}
}
