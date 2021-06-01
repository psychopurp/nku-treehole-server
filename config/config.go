package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Conf *viper.Viper
)

func Init(cfgDir string, isDev bool) {
	Conf = viper.New()
	Conf.AddConfigPath(cfgDir)
	if isDev {
		Conf.SetConfigName("config.dev")
	} else {
		Conf.SetConfigName("config.product")
	}
	Conf.SetConfigType("yml")
	if err := Conf.ReadInConfig(); err != nil {
		panic(err)
	}
	Conf.WatchConfig()
	Conf.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := Conf.ReadInConfig(); err != nil {
			panic(err)
		}
	})
}
