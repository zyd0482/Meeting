package logging

import (
    "os"
    "time"
    "fmt"
    "log"

    "meeting/pkg/setting"
)

func getLogFilePath() string {
    return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

func getLogFileFullPath() string {
    prefixPath := getLogFilePath()
    logSaveName := setting.AppSetting.LogSaveName
    logFileExt := setting.AppSetting.LogFileExt
    timeFormat := setting.AppSetting.TimeFormat;
    suffixPath := fmt.Sprintf("%s%s.%s", logSaveName, time.Now().Format(timeFormat), logFileExt)

    return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
    _, err := os.Stat(filePath)
    switch {
        case os.IsNotExist(err):
            mkDir()
        case os.IsPermission(err):
            log.Fatalf("Permission :%v", err)
    }

    handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Fail to OpenFile :%v", err)
    }

    return handle
}

func mkDir() {
    dir, _ := os.Getwd()
    err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
    if err != nil {
        panic(err)
    }
}