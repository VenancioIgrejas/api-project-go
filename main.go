package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VenancioIgrejas/api-project-go/pkg/setting"
	"github.com/VenancioIgrejas/api-project-go/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.AppSetting.Server.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.AppSetting.Server.ReadTimeout
	writeTimeout := setting.AppSetting.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.AppSetting.Server.Port)
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
}
