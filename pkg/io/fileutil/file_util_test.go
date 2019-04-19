package fileutil

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStringFromFile__SameString__Equals(t *testing.T) {
	file, _ := os.Create("test.txt")

	file.WriteString("123")

	str, _ := ReadStringFromFile("test.txt")

	assert.Equal(t, "123", str)

	os.Remove("test.txt")
}

func TestReadStringFromFile__DifferentStrings__NotEquals(t *testing.T) {
	file, _ := os.Create("test.txt")

	file.WriteString("123")

	str, _ := ReadStringFromFile("test.txt")

	assert.NotEqual(t, "321", str)

	os.Remove("test.txt")
}

func TestReadStringFromFile__FileDoesntExist__Error(t *testing.T) {
	_, err := ReadStringFromFile("test.txt")

	assert.Error(t, err)
}

func TestWriteStringToFile__SameString__Equals(t *testing.T) {
	WriteStringToFile("test.txt", "123")

	bytes, _ := ioutil.ReadFile("test.txt")

	assert.Equal(t, "123", string(bytes))

	os.Remove("test.txt")
}

func TestWriteStringToFile__DifferentStrings__NotEquals(t *testing.T) {
	WriteStringToFile("test.txt", "123")

	bytes, _ := ioutil.ReadFile("test.txt")

	assert.NotEqual(t, "321", string(bytes))

	os.Remove("test.txt")
}
