package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/damingerdai/blog-service/pkg/logger"

	"github.com/gin-gonic/gin"

	"github.com/damingerdai/blog-service/global"

	"github.com/damingerdai/blog-service/internal/model"
	"github.com/damingerdai/blog-service/internal/routers"
	"github.com/damingerdai/blog-service/pkg/setting"
)

func init() {
	var err error
	err = setupSetting()
	if err != nil {
		log.Fatalf("init setup Setting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setup DBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init setp Logger err: %v", err)
	}
}

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return nil
	}

	return err
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + os.PathSeparator + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
