import requests
import json

import datetime

times = 1000000


starttime = datetime.datetime.now()
for i in range(times):
    url = 'http://localhost:8080/api/v1/user/callWeatherRe'
    s = json.dumps({"cityIds": [2150126, 2150106, 2150096]})
    r = requests.post(url, data=s)

endtime = datetime.datetime.now()
print(endtime - starttime,"with redis", str(times) + "times")



starttime = datetime.datetime.now()
for i in range(times):
    url = 'http://localhost:8080/api/v1/user/callWeather'
    s = json.dumps({"cityIds": [2150126, 2150106, 2150096]})
    r = requests.post(url, data=s)

endtime = datetime.datetime.now()
print(endtime - starttime,"without redis",  str(times) + "times")