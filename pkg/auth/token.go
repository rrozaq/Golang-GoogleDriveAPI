package auth

import (
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"path"
)

const tokenFile = "token.json"

func WriteTokenFile(token *oauth2.Token) error {
	tokenBody, err := json.Marshal(token)
	if err != nil {
		return errors.New("marshaling authorization token failed")
	}

	fileDir, _ := os.Getwd()
	filePath := path.Join(fileDir, tokenFile)
	err = ioutil.WriteFile(filePath, tokenBody, 0644)

	if err != nil {
		return errors.New("writing token file failed")
	}
	return nil
}

func ReadTokenFile() (*oauth2.Token, error) {
	fileDir, _ := os.Getwd()
	filePath := path.Join(fileDir, tokenFile)
	tokenBody, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("reading token file failed")
	}

	var token oauth2.Token
	err = json.Unmarshal(tokenBody, &token)
	if err != nil {
		return nil, errors.New("unmarshaling token failed")
	}

	return &token, nil
}
