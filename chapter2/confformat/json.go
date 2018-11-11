package confformat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONData is our common data struct
// with JSON struct tags
type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ToJSON dumps the JSONData struct to
// a JSON format bytes.Buffer
func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode will decode into JSONData
func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// OtherJSONExamples shows ways to use types
// beyond structs and other useful functions
func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}

	fmt.Println("We can unmarshal into a map instead of a struct:", res)

	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("we can also use decoders/encoders to work with streams:", res)

	return nil
}
