package initialize

import (
	"fmt"
	"time"

	"github.com/user_service/pkg/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(config settings.Config) *gorm.DB {
	m := config.Databases
	var s = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic("Failed to connected to database")
	}
	SetPool(db, config)
	return db
}

func SetPool(db *gorm.DB, config settings.Config) {
	m := config.Databases
	sqlDb, err := db.DB()
	if err != nil {
		panic(err.Error)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)

}
