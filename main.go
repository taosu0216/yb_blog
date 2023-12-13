package main

import (
	"blog/global"
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg/setting"
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
	log.Println("init success")
}
func main() {
	router := routers.Router()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
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
