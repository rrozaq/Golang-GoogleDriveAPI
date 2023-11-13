package auth

import (
	"errors"
	"os"
)

var credentials = "credentials.json"

func readCredential() ([]byte, error) {
	readFile, err := os.ReadFile(credentials)
	if err != nil {
		return nil, errors.New("failed to read credentials")
	}

	return readFile, nil
}
