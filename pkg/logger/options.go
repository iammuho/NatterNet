package logger

// Option is the func interface to assign options
type Option func(*LoggerOptions)

type LoggerOptions struct {
	Level string
	Name  string
}

// WithLoggerLevel sets the logger level
func WithLoggerLevel(level string) Option {
	return func(o *LoggerOptions) {
		o.Level = level
	}
}

// WithLoggerName sets the logger name
func WithLoggerName(name string) Option {
	return func(o *LoggerOptions) {
		o.Name = name
	}
}
