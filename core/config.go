package core

import (
	"bytes"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() (err error) {

	box := packr.New("config", "./config")
	defaultConfig := box.Bytes("config.toml")

	configType := "toml"
	defaultPath := "./config"
	v := viper.New()
	// 读取默认配置 dev
	v.SetConfigName("config.toml")
	// v.AddConfigPath(defaultPath)
	v.SetConfigType(configType)
	// err = v.ReadInConfig()
	err = v.ReadConfig(bytes.NewReader(defaultConfig))
	if err != nil {
		panic(err)
	}

	configs := v.AllSettings()
	// 讲默认配置写入
	for k, v := range configs {
		viper.SetDefault(k, v)
	}

	env := os.Getenv("GO_ENV")
	// 根据配置的env 读取相应的配置信息
	if env != "" {
		viper.SetConfigName("config-" + env)
		viper.AddConfigPath(defaultPath)
		viper.SetConfigType(configType)
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
	}

	return err

}
