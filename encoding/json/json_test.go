package json

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock Data
type MockDataCar struct {
	Name string `json:"car"`
	Year int    `json:"year"`
}

func TestStrictUnMarshalJSON(t *testing.T) {

	assert := assert.New(t)

	// Define the test cases
	testCases := []struct {
		name     string
		data     []byte
		target   interface{}
		expected error
	}{
		{
			name:     "Success Case",
			data:     []byte(`{"car":"Ford","year":2021}`),
			target:   &MockDataCar{},
			expected: nil,
		},
		{
			name:     "Unknown Field",
			data:     []byte(`{"car":"Ford","year":2021,"color":"red"}`),
			target:   &MockDataCar{},
			expected: errors.New("Invalid JSON: json: unknown field \"color\""),
		},
		{
			name:     "Unexpected Data",
			data:     []byte(`{"car":"Ford","year":2021}{"car":"Ford","year":2021}`),
			target:   &MockDataCar{},
			expected: errors.New("unexpected data after JSON object"),
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := StrictUnMarshalJSON(tc.data, tc.target)

			if tc.expected != nil {
				assert.EqualError(err, tc.expected.Error())
			} else {
				assert.Nil(err)
			}

		})
	}

}
