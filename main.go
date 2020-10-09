package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

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

// @title 博客系统
// @version 1.0
// @description 大明二代使用Golang做项目
// @termsOfService https://github.com/damingerdai/blog-service-go
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
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

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
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	// setting.WatchSettingChange()
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return nil
	}

	if global.ServerSetting.RunMode == "debug" {
		global.DBEngine.LogMode(true)
	}
	return err
}

func setupLogger() error {
	filename := global.AppSetting.LogSavePath + string(os.PathSeparator) + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	fmt.Println(filename)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + string(os.PathSeparator) + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
