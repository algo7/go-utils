package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// StrictUnMarshalJSON is a function that unmarshals JSON data into a target interface
// It will return an error if there are any unknown fields in the JSON data and if there is any remaining data after decoding
// Example:
// data := []byte(`{"name": "John", "age": 30}`)
//
//	var person struct {
//	    Name string `json:"name"`
//	}
//
// err := StrictUnMarshalJSON(data, &person)
//
//	if err != nil {
//	    fmt.Println(err)
//	}
func StrictUnMarshalJSON(data []byte, target interface{}) error {

	// Create a new JSON decoder and disallow unknown fields
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	// Decode the JSON data into the target interface
	err := decoder.Decode(target)

	// Check if there was an error during decoding
	if err != nil {
		return fmt.Errorf("Invalid JSON: %s", err)
	}

	// Check if there is any remaining data in the decoder after decoding
	if decoder.More() {
		return errors.New("unexpected data after JSON object")
	}

	return nil
}
