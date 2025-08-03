package utils

import (
	"errors"
	"mime"
	"os"
	"strings"
)

func init() {
	mime.AddExtensionType(".md", "text/markdown")
}

func ReadSourceFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

func WriteToDestinationFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func VerifySourceFileType(path string) (bool, error) {
	splitStr := strings.Split(path, ".")
	if len(splitStr) == 0 {
		return false, errors.New("filename is empty")
	}

	if strings.ToLower(splitStr[len(splitStr)-1]) != "md" {
		return false, errors.New("only markdown (.md) file is supported as source")
	}

	return true, nil
}
