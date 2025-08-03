package utils

import (
	"os"
)

func ReadSourceFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

func WriteToDestinationFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
