package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Parse(cfg string) error {
	viper.SetDefault("listen", "127.0.0.1:8080")
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.path", "./app.log")
	viper.SetDefault("mysql.idle", 4)
	viper.SetDefault("mysql.limit", 100)
	viper.SetDefault("mysql.timeout", 1000)

	viper.SetConfigFile(cfg)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(errors.New("config file not found"))
		} else {
			fmt.Println(errors.New("config file was found but another error was produced"))
		}
		return err
	}

	fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	return nil
}
