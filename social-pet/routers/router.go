package routers

import (
	"social-pet/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/draw",
			beego.NSRouter("/add", &controllers.DrawController{}, "post:Draw"), //AI作画
			beego.NSRouter("/get", &controllers.DrawController{}, "post:Get"),  //AI作画信息
		),

		beego.NSNamespace("/upload",
			beego.NSRouter("/file", &controllers.UploadController{}, "post:UploadFile"),   //单文件上传
			beego.NSRouter("/files", &controllers.UploadController{}, "post:UploadFiles"), //多文件上传
			beego.NSRouter("/json", &controllers.UploadController{}, "post:UploadJson"),   //元数据JSON上传
		),

		beego.NSNamespace("/nft",
			beego.NSRouter("/mint", &controllers.NFTController{}, "post:Mint"),   //铸造
			beego.NSRouter("/count", &controllers.NFTController{}, "post:Count"), //铸造数量
		),

		beego.NSNamespace("/community",
			beego.NSRouter("/list", &controllers.CommunityController{}, "post:List"),                //动态列表
			beego.NSRouter("/publish", &controllers.CommunityController{}, "post:Publish"),          //自发/转发动态
			beego.NSRouter("/like", &controllers.CommunityController{}, "post:Like"),                //点赞/取消点赞
			beego.NSRouter("/comment/add", &controllers.CommunityController{}, "post:CommentAdd"),   //添加评论
			beego.NSRouter("/comment/list", &controllers.CommunityController{}, "post:CommentList"), //评论列表
		),

		beego.NSNamespace("/telegram",
			beego.NSRouter("/list", &controllers.TelegramController{}, "post:List"),          //群组列表
			beego.NSRouter("/list/user", &controllers.TelegramController{}, "post:UserList"), //群内成员列表
		),

		beego.NSNamespace("/message",
			beego.NSRouter("/send", &controllers.MessageController{}, "post:Send"), //发送消息
			beego.NSRouter("/list", &controllers.MessageController{}, "post:List"), //聊天列表
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
}
