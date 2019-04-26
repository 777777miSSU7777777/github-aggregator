// Package fileutil implements utility functions for working with files.
package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadStringFromFile reads file from specified path
// and returns its string content.
// If iotuil.ReadFile occurs any error, this will be returned.
func ReadStringFromFile(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(filePath))

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// WriteStringToFile creates new file with specified path and writes string to it.
func WriteStringToFile(filePath string, str string) error {
	file, err := os.Create(filePath)

	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(str)

	if err != nil {
		return err
	}

	return nil
}
