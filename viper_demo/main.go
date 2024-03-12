package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("read config err:", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unmarshal err:", err)
	}
	fmt.Println(config)
	fmt.Println(config.Mysql.Password)
}

type Config struct {
	Host  string
	Port  int
	Mysql Mysql
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
}
