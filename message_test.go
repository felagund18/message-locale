package message_locale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	LoadResource("messages", "en")

	assert.Equal(t, "Login by email", Format("auth.login"))
	assert.Equal(t, "The email a@a.com you have provided is not in our database",
		Format("auth.login_failed", "a@a.com"))
	assert.Equal(t, "Your email is a@a.com, and your username is a",
		Format("auth.info", "a", "a@a.com"))
}
