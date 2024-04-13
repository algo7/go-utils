package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {

	assert := assert.New(t)

	// Test case 1: File exists
	filename := "files.go"
	expected := true
	result := FileExists(filename)
	assert.Equal(expected, result, "Expected FileExists(%s) to be %v, but got %v", filename, expected, result)

	// Test case 2: File does not exist
	filename = "non-existent-file"
	expected = false
	result = FileExists(filename)
	assert.Equal(expected, result, "Expected FileExists(%s) to be %v, but got %v", filename, expected, result)

	// Test case 3: File exists but it is a directory
	filename = "."
	expected = false
	result = FileExists(filename)
	assert.Equal(expected, result, "Expected FileExists(%s) to be %v, but got %v", filename, expected, result)
}
