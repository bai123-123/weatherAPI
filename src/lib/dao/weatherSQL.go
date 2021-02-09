package dao

import (
	"fmt"
	"weatherAPI/src/define/jsonModel"
	"weatherAPI/src/lib/models"
	"weatherAPI/src/lib/mysql"
)

func InsertIntoWeather(op jsonModel.OpenWeatherResponse, sheetNum int) error {

	var sheetName = "`weather_info_current`"

	switch sheetNum {
	case 0:
		sheetName = "`weather_info_current`"
	case 1:
		sheetName = "`weather_info_d1`"
	case 2:
		sheetName = "`weather_info_d2`"
	case 3:
		sheetName = "`weather_info_d3`"

	}
	fmt.Println(sheetName)
	insertSql := "REPLACE INTO" + sheetName + "(`city_id`, `current_temperature`,`weather_condition`, `weather_code`, `highest_t_p`, `lowest_t_p`,`d_t`) VALUES(?,?,?,?,?,?,?)"

	if _, err := mysql.GetDb().Exec(insertSql, op.ID, op.Main.Temp, op.Weather[0].Main, op.Weather[0].ID, op.Main.TempMax, op.Main.TempMin, op.Dt); err != nil {
		return err
	}
	return nil
}

func QueryWeather(id int, sheetNum int) models.WeatherInfo{

	var sheetName = "`weather_info_current`"

	switch sheetNum {
	case 0:
		sheetName = "`weather_info_current`"
	case 1:
		sheetName = "`weather_info_d1`"
	case 2:
		sheetName = "`weather_info_d2`"
	case 3:
		sheetName = "`weather_info_d3`"
	}
	conn := mysql.GetDb()
	var CityId int
	var CurrentTemperature float64
	var WeatherCondition string
	var WeatherCode int
	var HighestTP float64
	var LowestTP float64
	var DT int
	sSql := "SELECT `*` FROM " + sheetName + "WHERE `city_id`=?"
	if err := conn.QueryRow(sSql, id).Scan(&CityId,&CurrentTemperature,&WeatherCondition,&WeatherCode,&HighestTP,&LowestTP,&DT); err != nil {
		fmt.Println(err)
	}
	return models.WeatherInfo{
		CityId:             CityId,
		CurrentTemperature: CurrentTemperature,
		WeatherCondition:   WeatherCondition,
		WeatherCode:        WeatherCode,
		HighestTP:          HighestTP,
		LowestTP:           LowestTP,
		DT:                 DT,
	}


}

