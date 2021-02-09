package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"weatherAPI/src/define/jsonModel"
	"weatherAPI/src/lib/dao"
	"weatherAPI/src/lib/log"

	//"weatherAPI/src/lib/log"

	"net/http"

	"strconv"
	"time"

	"weatherAPI/src/define/requestModels"
	"weatherAPI/src/define/responseModels"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong. %+v", time.Now().UTC().String())
}

func Nonono(c *gin.Context) {
	c.String(http.StatusNotFound, "Not Found. %+v", time.Now().UTC().String())
}

func SyncWeather(c *gin.Context) {
	var stat = new(requestModels.AsyncWeather)
	if err := c.ShouldBindJSON(&stat); err != nil {
		c.String(http.StatusBadRequest, "Wrong Body.")
		return
	} else {
		if err := handleCityRequest(stat.CityId); err != nil {
			log.WithContext(c).Errorf("fail to dao.CountContribution(), err: %v || adr: %s", err)
			c.Status(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, responseModels.CallWeatherR{
			stat.CityId,
		})
	}
}

func handleCityRequest(id int) error {

	for i := 0; i < 4; i++ {
		fmt.Println(i)
		now := time.Now()
		timePeriod, _ := time.ParseDuration("-" + strconv.Itoa(24*i) + "h")
		dt := now.Add(timePeriod).Unix()
		op, err := quryOpenWeather(id, dt)
		if err != nil {
			return err
		}


		if err := dao.InsertIntoWeather(op, i); err != nil {
			return err
		}

	}

	return nil

	//openWeatherResponse, err := quryOpenWeather(id, dt)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////fmt.Println(openWeatherResponse.Main.Temp)
	//fmt.Println(openWeatherResponse.Dt, dt)

	//singleTaskUpdateWeather(id, 0)
	//
	//singleTaskUpdateWeather(id, 1)
	//
	//singleTaskUpdateWeather(id, 2)
	//
	//singleTaskUpdateWeather(id, 3)

}

//func singleTaskUpdateWeather(id int64, index int) error {
//
//	now := time.Now()
//	oneDay, _ := time.ParseDuration("-24h")
//	twoDay, _ := time.ParseDuration("-48h")
//	threeDay, _ := time.ParseDuration("-72h")
//
//	mb := dao.MemberDao{mysql.Dbengine}
//	switch index {
//	case 0:
//		wd0, err := quryOpenWeather(id, now.Unix())
//		weather := models.WeatherInfoCurrent{
//			CityId:             wd0.ID,
//			CurrentTemperature: wd0.Main.Temp,
//			WeatherCondition:   wd0.Weather[0].Main,
//			WeatherCode:        wd0.Weather[0].ID,
//			HighestTP:          wd0.Main.TempMax,
//			LowestTP:           wd0.Main.TempMin,
//			DT:                 wd0.Dt,
//		}
//		mb.UpdateWeatherD0(weather)
//
//		if err != nil {
//			return err
//		}
//
//	case 1:
//
//		wd1, err := quryOpenWeather(id, now.Add(oneDay).Unix())
//
//		weather := models.WeatherInfoD1{
//			CityId:             wd1.ID,
//			CurrentTemperature: wd1.Main.Temp,
//			WeatherCondition:   wd1.Weather[0].Main,
//			WeatherCode:        wd1.Weather[0].ID,
//			HighestTP:          wd1.Main.TempMax,
//			LowestTP:           wd1.Main.TempMin,
//			DT:                 wd1.Dt,
//		}
//		mb.UpdateWeatherD1(weather)
//
//		if err != nil {
//			return err
//		}
//	case 2:
//		dt := now.Add(twoDay).Unix()
//		wd, err := quryOpenWeather(id, dt)
//		weather := models.WeatherInfoD2{
//			CityId:             wd.ID,
//			CurrentTemperature: wd.Main.Temp,
//			WeatherCondition:   wd.Weather[0].Main,
//			WeatherCode:        wd.Weather[0].ID,
//			HighestTP:          wd.Main.TempMax,
//			LowestTP:           wd.Main.TempMin,
//			DT:                 wd.Dt,
//		}
//		mb.UpdateWeatherD2(weather)
//
//		if err != nil {
//			return err
//		}
//
//	case 3:
//		wd, err := quryOpenWeather(id, now.Add(threeDay).Unix())
//		weather := models.WeatherInfoD3{
//			CityId:             wd.ID,
//			CurrentTemperature: wd.Main.Temp,
//			WeatherCondition:   wd.Weather[0].Main,
//			WeatherCode:        wd.Weather[0].ID,
//			HighestTP:          wd.Main.TempMax,
//			LowestTP:           wd.Main.TempMin,
//			DT:                 wd.Dt,
//		}
//		mb.UpdateWeatherD3(weather)
//
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
func quryOpenWeather(id int, dt int64) (jsonModel.OpenWeatherResponse, error) {

	var op jsonModel.OpenWeatherResponse
	params := url.Values{}
	Url, err := url.Parse("http://api.openweathermap.org/data/2.5/weather")
	if err != nil {
		return op, err
	}
	params.Set("id", strconv.Itoa(id))
	params.Set("exclude", string("daily"))
	params.Set("dt", strconv.FormatInt(dt, 10))
	params.Set("APPID", string("5013a32cc734deec3c8fefbbef1fab5f"))

	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return op, err
	}

	errPJ := json.Unmarshal(s, &op)

	if err != nil {
		fmt.Println(errPJ.Error())
	}
	return op, nil
}
