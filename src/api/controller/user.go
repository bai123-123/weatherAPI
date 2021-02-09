package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"strconv"
	"weatherAPI/src/lib/myredis"

	"weatherAPI/src/lib/dao"
	"weatherAPI/src/lib/models"

	"net/http"
	"weatherAPI/src/define/requestModels"

	//"weatherAPI/src/define/responseModels"

	//"weatherAPI/src/lib/models"

)

const (
	UserIdInURL = "adr"
)

func CallWeather(c *gin.Context) {

	fmt.Println(c.FullPath())

	var stat = new(requestModels.CallWeatherC)
	if err := c.ShouldBindJSON(&stat); err != nil {
		c.String(http.StatusBadRequest, "Wrong Body.")
		return
	} else {
		DataWeather := make(map[int][]models.WeatherInfo)

		for _, id := range stat.CityIds {
			s := make([] models.WeatherInfo, 0)
			for i := 0; i < 4; i++ {

				res := dao.QueryWeather(id, i)

				s = append(s, res)
			}
			DataWeather[id] = s

		}

		//var jsResponse = [...]string{ string(w0Js), w0}
		//c.PureJSON(200, w0)

		c.JSON(http.StatusOK, DataWeather)
	}

}

func CallWeatherWithRedis(c *gin.Context) {

	var stat = new(requestModels.CallWeatherC)
	if err := c.ShouldBindJSON(&stat); err != nil {
		c.String(http.StatusBadRequest, "Wrong Body.")
		return
	} else {
		DataWeather := make(map[int][]models.WeatherInfo)
		conn := myredis.RedisDefaultPool.Get()
		defer conn.Close()
		for _, id := range stat.CityIds {
			s := make([] models.WeatherInfo, 0)
			redisKey := "WC" + strconv.Itoa(id)
			ret, err := redis.Bytes(conn.Do("get", redisKey))
			weatherObj := []models.WeatherInfo{}
			if err != nil {
				for i := 0; i < 4; i++ {
					res := dao.QueryWeather(id, i)
					s = append(s, res)
				}
				retDate, _ := ffjson.Marshal(s)
				conn.Do("setex", redisKey, 2000, retDate)

				DataWeather[id] = s

			}else {
				ffjson.Unmarshal(ret, &weatherObj)

				DataWeather[id] = weatherObj
			}



		}

		c.JSON(http.StatusOK, DataWeather)
	}
}
