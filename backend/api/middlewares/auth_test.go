package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadBasicAccounts_Empty(t *testing.T) {
	allowedAccounts := LoadBasicAccounts()

	assert.Equal(t, gin.Accounts{}, allowedAccounts)
}

func TestLoadBasicAccounts(t *testing.T) {
	os.Setenv("ALLOWED_ACCOUNTS", "user:password;user2:password2;user3:password3;")
	allowedAccounts := LoadBasicAccounts()

	assert.Equal(t, gin.Accounts{
		"user":  "password",
		"user2": "password2",
		"user3": "password3",
	}, allowedAccounts)
}
