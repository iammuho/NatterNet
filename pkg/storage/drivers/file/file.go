package file

import (
	"os"

	"github.com/iammuho/natternet/pkg/storage/drivers"
)

type file struct{}

// NewFileStorage returns a new file storage
func NewFileStorage() drivers.DriverContext {
	return &file{}
}

// Get returns a file
func (f *file) Get() {

}

// Put puts a file
func (f *file) Put() {

}

// Delete deletes a file
func (f *file) Delete() {

}

// List lists files
func (f *file) List(path string) ([]string, error) {
	// list all files in a directory
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// log
	defer dir.Close()

	// get the list of files
	files, err := dir.Readdir(0)

	// log
	if err != nil {
		return nil, err
	}

	// prepare the list of files
	var fileList []string

	// loop through the files
	for _, file := range files {
		// append the file name to the list
		fileList = append(fileList, file.Name())
	}

	// return the list of files
	return fileList, nil
}
