package auth

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"net/http"
)

func Authenticate() (string, error) {
	config, err := ClientConfig()
	if err != nil {
		return "", errors.New("failed to get client")
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return "Open the following link in your browser, authenticate account and copy the authorization code: " + authURL, nil
}

func Client() (*http.Client, error) {
	token, err := ReadTokenFile()
	if err != nil {
		return nil, err
	}

	clientConfig, err := ClientConfig()
	if err != nil {
		return nil, err
	}

	return clientConfig.Client(context.Background(), token), nil

}

func ClientConfig() (*oauth2.Config, error) {
	getCredential, err := readCredential()
	if err != nil {
		return nil, err
	}

	scope := drive.DriveScope
	config, err := google.ConfigFromJSON(getCredential, scope)
	if err != nil {
		return nil, errors.New("parsing config file failed")
	}

	return config, nil
}
