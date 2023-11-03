package utils

import (
	"encoding/json"
	"os"
)

// TODO: Allow path to be passed in, or overridden for different environments

func ReadConfig[T any]() (T, error) {
	return readConfig[T]("./settings.json")
}

func readConfig[T any](path string) (T, error) {
	var val T
	data, err := os.ReadFile(path)
	if err != nil {
		return val, err
	}

	if err := json.Unmarshal(data, &val); err != nil {
		return val, err
	}
	return val, nil
}
