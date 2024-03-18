package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func RedirectRules() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求域名
		domain := c.Request.Host
		if domain != "novel.acgh.top" {
			log.Printf("host is %s\n", domain)
			log.Printf("will redirect %s\n", "https://novel.acgh.top"+c.Request.URL.String())
			c.Redirect(301, "https://novel.acgh.top"+c.Request.URL.String())
		}
		c.Next()
	}
}
