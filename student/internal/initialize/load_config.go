package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"student/global"
)

func LoadConfig() {
	vp := viper.New()
	vp.AddConfigPath("./config/")
	vp.SetConfigName("local")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := vp.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}
}
