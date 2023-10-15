package config

import (
	"gopkg.in/ini.v1"
	"log"
)

const (
	ConfPath = "./conf/db.ini"
)

var Conf *Config

func InitConfig() {
	cfg, err := ini.Load(ConfPath)
	if err != nil {
		panic(err.(any))
	}
	// ...
	conf := new(Config) //初始化一个结构体，返回指向他的指针
	err = cfg.MapTo(conf)
	if err != nil {
		panic(err.(any))
	}
	Conf = conf
	log.Println("database conf: ", Conf)
}
