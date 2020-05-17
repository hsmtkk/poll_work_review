package cred

import (
	"fmt"
	"os"
)

type CredentialLoader interface {
	LoadCredential() (string, string, error)
}

func New() CredentialLoader {
	return &impl{}
}

type impl struct{}

const UserNameKey = "TA_USER_NAME"
const PasswordKey = "TA_PASSWORD"

func (i *impl) LoadCredential() (string, string, error) {
	userName := os.Getenv(UserNameKey)
	if userName == "" {
		return "", "", fmt.Errorf("environment variable %s is not defined", UserNameKey)
	}
	password := os.Getenv(PasswordKey)
	if password == "" {
		return "", "", fmt.Errorf("environment variable %s is not defined", PasswordKey)
	}
	return userName, password, nil
}
