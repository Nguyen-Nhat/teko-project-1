package initialize

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"library/global"
)

func InitDatabase() {
	config := global.Config.Database
	dsn := "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	s := fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := sql.Open("postgres", s)
	if err != nil {
		global.Logger.Error("Init postgres error")
		panic(err)
	}
	if err := db.Ping(); err != nil {
		global.Logger.Error("Ping postgres error")
		panic(err)
	}
	global.Logger.Info("Init postgres success")
	global.Db = db
}
