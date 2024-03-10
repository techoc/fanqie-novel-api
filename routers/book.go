package routers

import (
	"github.com/gin-gonic/gin"
)

type BookRouter struct{}

func (s *BookRouter) InitBookRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	bookRouter := Router.Group("book")
	{
		bookRouter.GET("/search", bookApi.SearchBook)
		bookRouter.GET("/directory", bookApi.GetDirectoryByBookId)
		bookRouter.GET("/detail", bookApi.GetDetailByBookId)
	}
	return bookRouter
}
