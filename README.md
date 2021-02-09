# weatherAPI
Gin


API USAGE

POST WITH Json

Query weather with cityNumber list

http://localhost:8080/api/v1/user/callWeather

{
    "cityIds" : [2150126, 2150106, 2150096]
}


http://localhost:8080/api/v1/user/callWeatherRe

{
    "cityIds" : [2150126, 2150106, 2150096]
}

async weather with openMap API

http://localhost:8080/api/v1/asyncWeather

{
    "cityId" : 2150096

}
