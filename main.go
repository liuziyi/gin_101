package main

import (
	"fmt"
	"gin_101/pkg/setting"
	"gin_101/routers"
	"github.com/astaxie/beego/logs"
	"net/http"
)

func init() {
	setting.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ConfigSetting.ServerConfig.HttpPort)
	readTimeout := setting.ConfigSetting.ServerConfig.ReadTimeout
	writeTimeout := setting.ConfigSetting.ServerConfig.WriteTimeout
	server := http.Server{
		Addr:         endPoint,
		Handler:      routersInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	logs.Info("port: %v", endPoint)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
