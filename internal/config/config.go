package config

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
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
type Redis struct {
	Addr     string `json:"host"`
	Port     string `json:"port"`
	Database int    `json:"database"`
}
