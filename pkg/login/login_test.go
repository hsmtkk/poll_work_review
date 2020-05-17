package login_test

import (
	"fmt"
	"testing"

	"github.com/hsmtkk/poll_work_review/pkg/cred"
	"github.com/hsmtkk/poll_work_review/pkg/login"
	"github.com/stretchr/testify/assert"
)

func TestLocal(t *testing.T) {
}

func TestReal(t *testing.T) {
	userName, password, err := cred.New().LoadCredential()
	assert.Nil(t, err, "should be nil")
	cookie, err := login.New().GetAuthCookie(userName, password)
	assert.Nil(t, err, "should be nil")
	fmt.Println(cookie)
}
