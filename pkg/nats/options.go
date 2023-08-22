package nats

// Option is the func interface to assign options
type Option func(*NatsOptions)

// NatsOptions defines the NATS configuration
type NatsOptions struct {
	URL string
}

// WithNatsURL defines the NATS URL
func WithNatsURL(url string) Option {
	return func(o *NatsOptions) {
		o.URL = url
	}
}
