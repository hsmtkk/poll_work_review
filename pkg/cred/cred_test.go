package cred_test

import (
	"os"
	"testing"

	"github.com/hsmtkk/poll_work_review/pkg/cred"
	"github.com/stretchr/testify/assert"
)

func TestLoadCredential(t *testing.T) {
	loader := cred.New()
	reset()
	_, _, err := loader.LoadCredential()
	assert.Error(t, err, "error should happen")
	os.Setenv(cred.UserNameKey, "alpha")
	os.Setenv(cred.PasswordKey, "bravo")
	userName, password, err := loader.LoadCredential()
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, "alpha", userName, "should be equal")
	assert.Equal(t, "bravo", password, "should be equal")
	reset()
}

func reset() {
	os.Setenv(cred.UserNameKey, "")
	os.Setenv(cred.PasswordKey, "")
}
