package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MysqlConf struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     string
	}
}

func (c *Config) GetMysqlDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=Ture&loc=Asia/Shanghai",
		c.MysqlConf.User, c.MysqlConf.Password, c.MysqlConf.Host, c.MysqlConf.Port, c.MysqlConf.DBName)
	return dsn
}
