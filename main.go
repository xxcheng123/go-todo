package main

import (
	"fmt"
	"go-todo/conf"
	"go-todo/dao"
	"go-todo/models"
	"go-todo/routers"
)

const defaultConfFile = "./conf/config.ini"

func main() {
	conf.Init(defaultConfFile)
	//连接数据库
	dao.InitMySQL(conf.Conf.MySQLConfig)
	//绑定模型
	r := routers.SetupRputer()
	dao.DB.AutoMigrate(&models.Todo{})
	r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
}
