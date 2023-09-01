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
func (f *file) Get(fileName string) ([]byte, error) {
	// open the file
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	// close the file
	defer file.Close()

	// get the file info
	fileInfo, err := file.Stat()

	if err != nil {
		return nil, err
	}

	// prepare the buffer
	buffer := make([]byte, fileInfo.Size())

	// read the file
	_, err = file.Read(buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// Put puts a file
func (f *file) Put(fileName string, content []byte) error {
	// create the file
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	// close the file
	defer file.Close()

	// write the content
	_, err = file.Write(content)

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a file
func (f *file) Delete(fileName string) error {
	// delete the file
	err := os.Remove(fileName)

	if err != nil {
		return err
	}

	return nil
}

// List lists files
func (f *file) List(path string) ([]string, error) {
	// list all files in a directory
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// close the directory
	defer dir.Close()

	// get the list of files
	files, err := dir.Readdir(0)

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
