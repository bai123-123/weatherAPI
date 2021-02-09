package models

type WeatherInfoCurrent struct {
	CityId             int   `xorm:"pk" json:"city_id"`
	CurrentTemperature float64  `xorm:"notnull" json:"current_temperature"`
	WeatherCondition   string  `xorm: "notnull" json:"weather_condition"`
	WeatherCode        int   `xorm: "notnull" json:"weather_code"`
	HighestTP          float64 `xorm:"notnull" json:"highest_tp"`
	LowestTP           float64 `xorm: "notnull" json:"lowest_tp"`
	DT                 int   `xorm:"updated" json:"dt"`
}

type WeatherInfoD1 struct {
	CityId             int   `xorm:"pk" json:"city_id"`
	CurrentTemperature float64  `xorm:"notnull" json:"current_temperature"`
	WeatherCondition   string  `xorm: "notnull" json:"weather_condition""`
	WeatherCode        int   `xorm: "notnull" json:"weather_code"`
	HighestTP          float64 `xorm:"notnull" json:"highest_tp"`
	LowestTP           float64 `xorm: "notnull" json:"lowest_tp"`
	DT                 int   `xorm:"notnull" json:"dt"`
}


type WeatherInfoD2 struct {
	CityId             int   `xorm:"pk" json:"city_id"`
	CurrentTemperature float64  `xorm:"notnull" json:"current_temperature"`
	WeatherCondition   string  `xorm: "notnull" json:"weather_condition""`
	WeatherCode        int   `xorm: "notnull" json:"weather_code"`
	HighestTP          float64 `xorm:"notnull" json:"highest_tp"`
	LowestTP           float64 `xorm: "notnull" json:"lowest_tp"`
	DT                 int   `xorm:"notnull" json:"dt"`
}


type WeatherInfoD3 struct {
	CityId             int   `xorm:"pk" json:"city_id"`
	CurrentTemperature float64  `xorm:"notnull" json:"current_temperature"`
	WeatherCondition   string  `xorm: "notnull" json:"weather_condition""`
	WeatherCode        int   `xorm: "notnull" json:"weather_code"`
	HighestTP          float64 `xorm:"notnull" json:"highest_tp"`
	LowestTP           float64 `xorm: "notnull" json:"lowest_tp"`
	DT                 int   `xorm:"notnull" json:"dt"`
}


type WeatherInfo struct {
	CityId             int   `json:"city_id"`
	CurrentTemperature float64  `json:"current_temperature"`
	WeatherCondition   string  `json:"weather_condition"`
	WeatherCode        int   `json:"weather_code"`
	HighestTP          float64 `json:"highest_tp"`
	LowestTP           float64 `json:"lowest_tp"`
	DT                 int   `json:"dt"`
}