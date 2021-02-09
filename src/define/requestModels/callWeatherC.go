package requestModels

type CallWeatherC struct {
	CityIds []int `form: "cityIds"`
}

type AsyncWeather struct {
	CityId int `form: "cityId"`
}
