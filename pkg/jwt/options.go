package jwt

// Option is the func interface to assign options
type Option func(*JwtOptions)

// JwtOptions defines the jwt configuration
type JwtOptions struct {
	PublicKeyPath  string
	PrivateKeyPath string
	Kid            string
	Issuer         string
	Subject        string
}

// WithJWTPublicKeyPath aims to set the public key path
func WithJWTPublicKeyPath(path string) Option {
	return func(o *JwtOptions) {
		o.PublicKeyPath = path
	}
}

// WithJWTPrivateKeyPath aims to set the private key path
func WithJWTPrivateKeyPath(path string) Option {
	return func(o *JwtOptions) {
		o.PrivateKeyPath = path
	}
}

// WithJWTIssuer aims to set the issuer
func WithJWTIssuer(issuer string) Option {
	return func(o *JwtOptions) {
		o.Issuer = issuer
	}
}

// WithJWTSubject aims to set the subject
func WithJWTSubject(subject string) Option {
	return func(o *JwtOptions) {
		o.Subject = subject
	}
}

// WithJWTKeyID aims to set the key id
func WithJWTKeyID(kid string) Option {
	return func(o *JwtOptions) {
		o.Kid = kid
	}
}
