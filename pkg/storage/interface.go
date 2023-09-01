package storage

import "github.com/iammuho/natternet/pkg/storage/drivers"

// StorageContext is the interface for the storage
type StorageContext interface {
	Driver() drivers.DriverContext
}
