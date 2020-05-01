package middlewares

import "github.com/gin-gonic/gin"

func JsonHeaders(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Accept", "application/json")

	c.Next()
}
