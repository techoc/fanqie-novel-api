package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RedirectRules() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求域名
		domain := c.Request.Host
		fmt.Printf("host is %s\n", domain)
		c.Next()
	}
}
