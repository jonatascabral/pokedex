package middlewares

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func LoadBasicAccounts() gin.Accounts {
	accounts := strings.Split(os.Getenv("ALLOWED_ACCOUNTS"), ";")
	allowedAccounts := make(gin.Accounts)
	if len(accounts) == 0 {
		return allowedAccounts
	}

	for _, account := range accounts {
		userPass := strings.Split(account, ":")
		if len(userPass) == 2 {
			allowedAccounts[userPass[0]] = userPass[1]
		}
	}

	return allowedAccounts
}

func AuthHandler(c *gin.Context) {
	gin.BasicAuth(LoadBasicAccounts())

	c.Next()
}
