package file

import (
	"log"
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
func (f *file) List(path string) {
	// list all files in a directory
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// log
	defer dir.Close()

	// get the list of files
	files, err := dir.Readdir(0)

	// log
	if err != nil {
		panic(err)
	}

	// log
	for _, file := range files {
		// log
		if file.IsDir() {
			// log
			log.Print("Directory: ", file.Name())
		} else {
			// log
			log.Print("File: ", file.Name())
		}
	}

}
