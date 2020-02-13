package main

import (
    "fmt"
    "net/http"
    "context"
    "log"
    "os"
    "os/signal"
    "time"


    "meeting/routers"
    "meeting/models"
    "meeting/pkg/setting"
    "meeting/pkg/logging"
)

func main() {
    setting.Setup()
    models.Setup()
    logging.Setup()

    router := routers.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
        Handler:        router,
        ReadTimeout:    setting.ServerSetting.ReadTimeout,
        WriteTimeout:   setting.ServerSetting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    go func() {
        if err := s.ListenAndServe(); err != nil {
            log.Printf("Listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <- quit

    log.Println("Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    log.Println("Server exiting")
}