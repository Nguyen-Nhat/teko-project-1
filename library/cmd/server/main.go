package main

import (
	"fmt"
	"library/global"
	"library/internal/initialize"
)

func main() {
	initialize.Run()
	global.Logger.Info(fmt.Sprint(global.Config))
}
