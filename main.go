package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-server-template/config"
	// "go-server-template/middleware/global/jwt"
	"go-server-template/cron"
	"go-server-template/middleware/global/log"
	"go-server-template/pkg/db"
	"go-server-template/pkg/redis"
	"go-server-template/routers"
	"net/http"
	"time"
)

func init() {
	projectConfig.InitConfig()
	DB.InitDBGorm()
	Redis.InitRedis()
	r := routers.InitRouter()
	r.Use(logMiddleware.LogerMiddleware())
	// r.Use(JWTMiddleware.JWT())
}

// @title 项目模版`
// @version 1.0`
// @description 为了自己以后方便写点小东西，搭个模版`
// @termsOfService [http://swagger.io/terms/](http://swagger.io/terms/)`
// @contact.name 黄日明`
// @contact.url [http://www.swagger.io/support](http://www.swagger.io/support)`
// @contact.email support@swagger.io`
// @license.name Apache 2.0`
// @license.url [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`
func main() {
	config := projectConfig.AppConfig
	gin.SetMode(config.ServerConfig.RUN_MODE)
	routersInit := routers.InitRouter()
	readTimeout := config.ServerConfig.READ_TIMEOUT
	writeTimeout := config.ServerConfig.WRITE_TIMEOUT
	endPoint := fmt.Sprintf(":%d", config.ServerConfig.HTTP_PORT)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	cronTask.CronTaskInit()

	log.Printf("开始监听端口 %s", endPoint)
	server.ListenAndServe()

}
