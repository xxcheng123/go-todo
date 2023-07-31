package dao

import (
	"fmt"
	"go-todo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// 初始化数据库
func InitMySQL(cfg *conf.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)
	DB, err = gorm.Open(mysql.Open(dsn))
	if err == nil {
		return
	}

	return err

}
