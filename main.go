package main

import (
	"blog/global"
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg/logger"
	"blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	var err error
	if err = setUpsetting(); err != nil {
		log.Fatalln("init err is : ", err)
	}
	if err = initDBEngine(); err != nil {
		log.Fatalln("init db err is : ", err)
	}
	if err = setLogger(); err != nil {
		log.Fatalln("init logger err is : ", err)
	}
	log.Println("init success")
}

// @title Yblue`s blog
// @version 1.0
// @description Go blog
// @termsOfService https://blog.yblue.top
func main() {
	router := routers.Router()
	s := &http.Server{
		Addr:           ":44444",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	s.ListenAndServe()
}
func setUpsetting() error {
	settings, err := setting.InitSetting()
	if err != nil {
		return err
	}
	err = settings.ReadSections("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSections("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSections("Mysql", &global.MysqlSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func initDBEngine() error {
	var err error
	global.MysqlEngine, err = model.NewDBEngine(global.MysqlSetting)
	if err != nil {
		return err
	}
	return nil
}
func setLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogSaveName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
