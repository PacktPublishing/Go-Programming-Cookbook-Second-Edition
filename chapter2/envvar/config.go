package envvar

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// LoadConfig will load files optionally from the json file stored
// at path, then will override those values based on the envconfig
// struct tags. The envPrefix is how we prefix our environment
// variables.
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from env")
}

// LoadFile unmarshalls a json file into a config struct
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}
