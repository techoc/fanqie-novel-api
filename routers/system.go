package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemRouter struct{}

func (s *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	//systemRouter := Router.Group("sys")
	systemRouter := Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return systemRouter
}
