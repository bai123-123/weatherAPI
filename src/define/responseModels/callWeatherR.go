package responseModels

type CallWeatherR struct {
	CityId int `form: "cityIds"`
}

type ResponseWeather struct {
	CityId             int     ` json:"city_id"`
	CurrentTemperature float64 `json:"current_temperature"`
	WeatherCondition   string  ` json:"weather_condition"`
	WeatherCode        int     `json:"weather_code"`
	HighestTP          float64 ` json:"highest_tp"`
	LowestTP           float64 `json:"lowest_tp"`
	DT                 int     ` json:"dt"`
}
