package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"

    "git.codepku.com/examinate/exam/routers"
    "git.codepku.com/examinate/exam/models"
    "git.codepku.com/examinate/exam/pkg/setting"
    "git.codepku.com/examinate/exam/pkg/logging"
    // "git.codepku.com/examinate/exam/pkg/util"
)

func init() {
    setting.Setup()
    models.Setup()
    logging.Setup()
    // util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An exam of gin
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

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