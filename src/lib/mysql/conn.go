package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"time"

	"sync"
	_const "weatherAPI/src/const"
	"weatherAPI/src/lib/log"
	//"weatherAPI/src/lib/models"
	//"gorm.io/gorm"
)

var (
	conn     *sql.DB
	lock     sync.Mutex
	dsn      string
	//Dbengine *gorm.DB
)

func init() {
	dsn = _const.DSN
}

// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func GetDb() *sql.DB {
	if conn != nil {
		return conn
	}
	lock.Lock()
	defer lock.Unlock()
	if conn != nil {
		return conn
	}
	if db, err := sql.Open("mysql", dsn); err != nil {
		log.Logger().Fatalf("can't init db, err: %v", err)
		panic(err)
	} else {
		if err := db.Ping(); err != nil {
			log.Logger().Fatalf("can't ping db, err: %v", err)
			panic(err)
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		conn = db
		return conn
	}
}



//func OrmEngine() (*gorm.DB, error) {
//	engine, err := xorm.NewEngine("mysql", dsn)
//
//	if err != nil {
//		return nil, err
//	}
//
//	//engine.ShowSQL(true)
//
//	if err := engine.Sync2(new(models.WeatherInfoCurrent)); err != nil {
//		return nil, err
//	}
//
//	if err := engine.Sync2(new(models.WeatherInfoD1)); err != nil {
//		return nil, err
//	}
//
//	if err := engine.Sync2(new(models.WeatherInfoD2)); err != nil {
//		return nil, err
//	}
//
//	if err := engine.Sync2(new(models.WeatherInfoD3)); err != nil {
//		return nil, err
//	}
//
//	orm := new(Orm)
//	orm.Engine = engine
//	Dbengine = orm
//
//	cacher := xrc.NewRedisCacher("localhost:6379", "", xrc.DEFAULT_EXPIRATION, engine.Logger())
//	engine.SetDefaultCacher(cacher)
//
//	return orm, nil
//
//}
