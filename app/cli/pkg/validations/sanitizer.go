package validations

import (
	"io"
	"strings"
)

func Sanitize(m string) string {
	return strings.TrimSpace(strings.ReplaceAll(m, "\\", ""))
}

func IsNilThenDefault(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}

	return *value
}

func IsNilThenEmpty(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func GetIOIntoBytes(ioBuffer io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(ioBuffer)
	if err != nil {
		return nil, err
	}

	return data, nil
}
