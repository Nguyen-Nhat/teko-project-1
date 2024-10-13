package initialize

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"student/global"
)

func InitDatabase() {
	config := global.Config.Database
	dsn := "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	s := fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		global.Logger.Error("Init postgres error")
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.Database.MaxIdleConn)
	sqlDB.SetMaxOpenConns(global.Config.Database.MaxOpenConn)
	global.Db = db
}
