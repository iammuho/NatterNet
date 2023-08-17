package http

import "time"

// Option is the func interface to assign options
type Option func(*HTTPServerOptions)

// HTTPServerOptions defines the http server configuration
type HTTPServerOptions struct {
	ServerHeader string
	AppName      string

	CaseSensitive bool
	StrictRouting bool
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	BodyLimit     int
}

// WithHTTPServerHeader defines the HTTPServer Header
func WithHTTPServerHeader(header string) Option {
	return func(o *HTTPServerOptions) {
		o.ServerHeader = header
	}
}

// WithHTTPServerAppName defines the HTTPServer App Name
func WithHTTPServerAppName(appName string) Option {
	return func(o *HTTPServerOptions) {
		o.AppName = appName
	}
}

// WithHTTPServerCaseSensitive defines the HTTPServer Case Sensitive
func WithHTTPServerCaseSensitive(isCaseSensitive bool) Option {
	return func(o *HTTPServerOptions) {
		o.CaseSensitive = isCaseSensitive
	}
}

// WithHTTPServerStrictRouting defines the HTTPServer Strict Routing
func WithHTTPServerStrictRouting(isStrictRouting bool) Option {
	return func(o *HTTPServerOptions) {
		o.StrictRouting = isStrictRouting
	}
}

// WithHTTPServerReadTimeout defines the HTTPServer Read Timeout
func WithHTTPServerReadTimeout(readTimeout time.Duration) Option {
	return func(o *HTTPServerOptions) {
		o.ReadTimeout = readTimeout
	}
}

// WithHTTPServerWriteTimeout defines the HTTPServer Write Timeout
func WithHTTPServerWriteTimeout(writeTimeout time.Duration) Option {
	return func(o *HTTPServerOptions) {
		o.WriteTimeout = writeTimeout
	}
}

// WithHTTPServerBodyLimit defines the HTTPServer Body Limit
func WithHTTPServerBodyLimit(bodyLimit int) Option {
	return func(o *HTTPServerOptions) {
		o.BodyLimit = bodyLimit
	}
}
