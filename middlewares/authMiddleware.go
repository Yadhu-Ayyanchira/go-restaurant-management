package middlewares

import "github.com/gin-gonic/gin"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Add token verification logic
		c.Next()
	}
}
