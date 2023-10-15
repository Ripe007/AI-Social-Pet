package service

import (
	"fmt"
	"log"
	"social-pet/config"
	"social-pet/global"
	"social-pet/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func InitDb() {
	var (
		err    error
		engine *xorm.Engine
	)

	//连接数据库
	engine, err = xorm.NewEngine(config.Conf.DriverName, GetDataSourceName())
	if err != nil {
		panic(err.(any))
	}
	//连接测试
	if err = engine.Ping(); err != nil {
		panic(err.(any))
		return
	}

	global.DB = engine

	log.Println("connect database success")

}

func GetDataSourceName() string {
	log.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8\n",
		config.Conf.User,
		config.Conf.Password,
		config.Conf.Ip,
		config.Conf.Port,
		config.Conf.Db)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.Conf.User,
		config.Conf.Password,
		config.Conf.Ip,
		config.Conf.Port,
		config.Conf.Db,
	)

}

func InitTables() {
	if err := global.DB.Sync2(
		new(models.AnimalInfo),           //物种表
		new(models.TelegramInfo),         //群组表
		new(models.MessageInfo),          //聊天消息表
		new(models.UserTelegramInfo),     //用户-群组关系表
		new(models.DrawInfo),             //作画表
		new(models.MintInfo),             //铸造记录表
		new(models.CommunityContextInfo), //社区动态内容表
		new(models.CommunityConfigInfo),  //社区动态配置表
		new(models.CommunityLikeInfo),    //社区动态点赞表
		new(models.CommunityCommentInfo), //社区动态评论表
	); err != nil {
		panic(err.(any))
	}
}
