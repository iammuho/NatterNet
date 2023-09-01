package drivers

// DriverContext is the interface for the storage driver
// TODO: add/refactor methods
type DriverContext interface {
	Get(string) ([]byte, error)
	Put(string, []byte) error
	Delete(string) error
	List(string) ([]string, error)
}
