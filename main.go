package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"

    "meeting/models"
    "meeting/pkg/logging"
    "meeting/pkg/setting"
    "meeting/routers"
    "meeting/pkg/util"
    "meeting/admin"
)

func init() {
    setting.Setup()
    models.Setup()
    logging.Setup()
    util.Setup()
}

func main() {
    gin.SetMode(setting.ServerSetting.RunMode)
    handler := gin.New()
    handler.InitRouter()
    handler.InitAdmin()

    readTimeout := setting.ServerSetting.ReadTimeout
    writeTimeout := setting.ServerSetting.WriteTimeout
    endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
    maxHeaderBytes := 1 << 20

    server := &http.Server{
        Addr:           endPoint,
        Handler:        handler,
        ReadTimeout:    readTimeout,
        WriteTimeout:   writeTimeout,
        MaxHeaderBytes: maxHeaderBytes,
    }

    log.Printf("[info] start http server listening %s", endPoint)

    server.ListenAndServe()
}
