package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/techoc/fanqie-novel-api/conf"
	"github.com/techoc/fanqie-novel-api/models"
	"github.com/techoc/fanqie-novel-api/pkg/global"
	"github.com/techoc/fanqie-novel-api/routers"
	"log"
	"net/http"
)

func init() {
	conf.Setup()
	models.Setup()
	//logging.Setup()
	//gredis.Setup()
	//util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(global.ServerConf.RunMode)

	readTimeout := global.ServerConf.ReadTimeout
	writeTimeout := global.ServerConf.WriteTimeout
	endPoint := fmt.Sprintf(":%d", global.ServerConf.HttpPort)
	maxHeaderBytes := 1 << 20

	Router := gin.New()
	Router.Use(gin.Recovery())
	//if gin.Mode() == gin.DebugMode {
	Router.Use(gin.Logger())
	//Router.Use(middleware.RedirectRules())
	//}
	router := routers.RouterGroupApp
	router.InitBookRouter(&Router.RouterGroup)
	router.InitChapterRouter(&Router.RouterGroup)
	router.InitSystemRouter(&Router.RouterGroup)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        Router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("[error] http server err: %v\n", err)
		return
	}

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
