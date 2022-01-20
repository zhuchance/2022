package main

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/common"
	"github.com/EDDYCJY/go-gin-example/controller"
	"github.com/EDDYCJY/go-gin-example/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)



func main() {
	//db := common.GetDB()
	db := common.InitDB()
	fmt.Println(db)
	r := gin.Default()
	//r = CollerRoute(r)
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Loing)
	r.GET("/api/auth/info",middleware.AuthMiddleware(), controller.Info)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务

}

func InitConfig()  {
	workDir,_ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir+"/config")
}

