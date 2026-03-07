package initialize

import (
	"fmt"
	"time"

	"github.com/user_service/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	m := global.Config.Databases
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic("Failed to connected to database")
	}
	SetPool(db)
	return db
}

func SetPool(db *gorm.DB) {
	m := global.Config.Databases
	sqlDb, err := db.DB()
	if err != nil {
		panic(err.Error)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)

}
