package initialize

import (
	"fmt"
	"student/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDatabase()
	r := InitRouter()
	_ = r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
}
