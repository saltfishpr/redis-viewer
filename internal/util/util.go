// @description:
// @file: util.go
// @date: 2022/02/07

// Package util
package util

import (
	"errors"
	"os"
)

// GetHomeDirectory returns the users home directory.
func GetHomeDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home, nil
}

// CreateDirectory creates a new directory given a name.
func CreateDirectory(name string) error {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteDirectory deletes a directory given a name.
func DeleteDirectory(name string) error {
	err := os.RemoveAll(name)

	return err
}
