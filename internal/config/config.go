package config

import "github.com/spf13/viper"

var Conf Config

type Config struct {
	Jwt   Jwt   `json:"jwt"`
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type Jwt struct {
	Key  string `json:"key"`
	Hour int    `json:"hour"`
}

type Mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
type Redis struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Database int    `json:"database"`
}

func Init() error {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	return nil
}
