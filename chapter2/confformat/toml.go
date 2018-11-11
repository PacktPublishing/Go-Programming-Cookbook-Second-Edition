package confformat

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

// TOMLData is our common data struct
// with TOML struct tags
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML dumps the TOMLData struct to
// a TOML format bytes.Buffer
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)

	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

// Decode will decode into TOMLData
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
