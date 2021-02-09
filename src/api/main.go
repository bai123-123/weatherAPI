package main

import (

	"log"
	"weatherAPI/src/api/controller"
	"weatherAPI/src/api/middleware"

	_ "weatherAPI/src/lib/myredis"
	"weatherAPI/src/lib/mysql"

	//"weatherAPI/src/api/middleware"
	//"weatherAPI/src/lib/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"time"
)

func init() {
	os.Setenv("TZ", "UTC")
}

func initEnv() {
	//// try to connect db, maybe cause panic
	_ = mysql.GetDb()
	//_, _ = mysql.OrmEngine()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	initEnv()


	r := getGinEngine()
	//getGinEngine()
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func getGinEngine() *gin.Engine {
	r := gin.New()
	r.Use(middleware.AccessLog)

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	setRouter(r)

	return r
}

func setRouter(r *gin.Engine) {
	r.NoRoute(controller.Nonono)
	v1api := r.Group("/api/v1")
	v1api.GET("/ping", controller.Ping)
	v1api.POST("/asyncWeather", controller.SyncWeather)

	user := v1api.Group("/user")
	{
		user.POST("/callWeather", controller.CallWeather)
		user.POST("/callWeatherRe", controller.CallWeatherWithRedis)
	}
}
