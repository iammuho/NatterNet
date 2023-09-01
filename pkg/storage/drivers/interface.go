package drivers

// DriverContext is the interface for the storage driver
// TODO: add/refactor methods
type DriverContext interface {
	Get()
	Put()
	Delete()
	List(string)
}
