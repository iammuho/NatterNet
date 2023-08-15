package http

// Option is the func interface to assign options
type Option func(*HTTPServerOptions)

// HTTPServerOptions defines the http server configuration
type HTTPServerOptions struct {
	Address      string
	Port         int
	IsTLSEnabled bool
}

// WithHTTPServerAddress defines the HTTPServer Listening Address
func WithHTTPServerAddress(address string) Option {
	return func(o *HTTPServerOptions) {
		o.Address = address
	}
}

// WithHTTPServerPort defines the HTTPServer Listening Port
func WithHTTPServerPort(port int) Option {
	return func(o *HTTPServerOptions) {
		o.Port = port
	}
}

// WithHTTPServerTLSEnabled defines the HTTPServer TLS Enabled
func WithHTTPServerTLSEnabled(isTLSEnabled bool) Option {
	return func(o *HTTPServerOptions) {
		o.IsTLSEnabled = isTLSEnabled
	}
}
