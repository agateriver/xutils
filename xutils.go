package xutils

import (
	"os"
	"errors"
	"fmt"
)

// IsFileExists if file exists, return true, otherwise false and err
func IsFileExists(file string) (bool, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	if fi != nil && fi.IsDir() {
		return false, errors.New(fmt.Sprintf("%v cannot be directory.", file))
	}
	return true, nil
}

// Noop is a tricky for keep unused variables
func Noop(v ...interface{}) {}
