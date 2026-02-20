package initialize

import (
	"fmt"
	"time"

	"github.com/auth_service/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	m := global.Config.Databases
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic("Failed to connected to database")
	}
	global.DB = db
	fmt.Println("Connected to database")
}

func SetPool() {
	m := global.Config.Databases
	sqlDb, err := global.DB.DB()
	if err != nil {
		panic(err.Error)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}
