package main

import (
	"social-pet/beego/cors"
	"social-pet/beego/handler"
	"social-pet/config"
	"social-pet/egret"
	"social-pet/global"
	_ "social-pet/routers"
	"social-pet/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//设置logger
	err := logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/logger.log", "separate":["error", "warning", "info"],"level":3}`)
	if err != nil {
		logs.Error(err)
	}

	//跨域
	cors.InitCors()

	//开启body copy
	beego.BConfig.CopyRequestBody = true

	//404自定义返回
	beego.ErrorHandler("404", handler.Handle404Error)

	//初始化结构体
	config.InitConfig()

	//初始化数据库
	service.InitDb()

	// 设置Mysql连接
	global.DB.SetMaxOpenConns(64)

	global.DB.SetMaxIdleConns(64)

	//同步表结构
	service.InitTables()

	//初始化web3，连接rpc
	service.InitWeb3()

	//初始化定时器
	egret.InitEgretJob()

	beego.Run()
}
