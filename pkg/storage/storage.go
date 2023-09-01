package storage

import (
	"github.com/iammuho/natternet/pkg/storage/drivers"
	"github.com/iammuho/natternet/pkg/storage/drivers/file"
)

type storage struct {
	driver  drivers.DriverContext
	options StorageOptions
}

func NewStorage(opts ...Option) (StorageContext, error) {
	// Setup the driver
	options := StorageOptions{}
	for _, o := range opts {
		o(&options)
	}

	switch options.Driver {
	case DriverFile:
		return &storage{
			driver:  file.NewFileStorage(),
			options: options,
		}, nil
	}

	return &storage{}, nil
}

func (s *storage) Driver() drivers.DriverContext {
	return s.driver
}
