import pymysql
import json

db = pymysql.connect(host='localhost', port=3306, user='root', password='root', database='weatherDB', )

cur = db.cursor()

cityDic = {}
with open("city.list.json", 'r') as load_f:
    cityDic = json.load(load_f)


print(len(cityDic))

data = []
for i in cityDic:
    data.append((i['id'], i['name'], i['state'], i['country'], i['coord']['lon'], i['coord']['lat']))


print(len(data))

# remit_ids = [('1234', 'CAD'), ('5678', 'HKD')]
#
try:
    cur.executemany("INSERT INTO cityID (id, nameCT, state,country,lon,lat) "
                    + " VALUES (%s,%s, %s, %s,%s,%s)", data)
    assert cur.rowcount == len(data), 'my error message'
    print(cur.rowcount)
    db.commit()
    print('xxxx')
except Exception as e:
    db.rollback()
    print(e)

finally:
    cur.close()
    db.close()

