package storage

// Option is the func interface to assign options
type Option func(*StorageOptions)

type Driver string

const (
	// DriverFile is the file driver
	DriverFile Driver = "file"
	// DriverAWS is the AWS driver
	DriverAWS Driver = "aws"
)

// StorageOptions defines the options for the storage
type StorageOptions struct {
	Driver Driver
}

// WithStorageDriver sets the storage driver
func WithStorageDriver(driver string) Option {
	return func(o *StorageOptions) {
		o.Driver = Driver(driver)
	}
}
