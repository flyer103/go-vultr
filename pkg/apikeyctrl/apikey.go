// API Key Controller.
package apikeyctrl

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ErrNoHomeEnv   = errors.New("No Home environment")
	ErrEmptyAPIKey = errors.New("API Key is empty")
)

type APIKeyController struct {
	FilePath string
}

func New() (*APIKeyController, error) {
	homePath := os.Getenv("HOME")
	if homePath == "" {
		return nil, ErrNoHomeEnv
	}

	vultrDir := filepath.Join(homePath, ".vultr")
	if err := os.MkdirAll(vultrDir, 0755); err != nil {
		return nil, err
	}

	credentialFile := filepath.Join(vultrDir, "credential")

	ac := &APIKeyController{
		FilePath: credentialFile,
	}
	return ac, nil
}

func (ac *APIKeyController) Store(apiKey string) error {
	if apiKey == "" {
		return ErrEmptyAPIKey
	}

	encryptData, err := Encrypt(EnctryptKey, []byte(apiKey))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(ac.FilePath, encryptData, 0755)
}

func (ac *APIKeyController) Update(apiKey string) error {
	return ac.Store(apiKey)
}

func (ac *APIKeyController) Get() (string, error) {
	dataBytes, err := ioutil.ReadFile(ac.FilePath)
	if err != nil {
		return "", err
	}

	apiKeyBytes, err := Decrypt(EnctryptKey, dataBytes)
	if err != nil {
		return "", err
	}

	return string(apiKeyBytes), nil
}

func (ac *APIKeyController) Delete() error {
	return os.Remove(ac.FilePath)
}
