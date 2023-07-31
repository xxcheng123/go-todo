package conf

import "gopkg.in/ini.v1"

var Conf = new(ApplicationConfig)

type ApplicationConfig struct {
	Port         uint16 `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

type MySQLConfig struct {
	User         string `ini:"user"`
	Password     string `ini:"password"`
	Host         string `ini:"host"`
	Port         uint16 `ini:"port"`
	DatabaseName string `ini:"dbName"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
