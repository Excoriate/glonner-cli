package common

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func GetEnv(key string) (string, error) {
	env := os.Getenv(key)
	if env == "" {
		return "", errors.New(fmt.Sprintf("environment variable %s is not set", key))
	}

	return env, nil
}
