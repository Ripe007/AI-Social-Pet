package handler

import (
	"net/http"
	"social-pet/msg"
	"social-pet/util"

	"github.com/astaxie/beego/logs"
)

// 404报错
var Handle404Error = func(writer http.ResponseWriter, request *http.Request) {
	data := msg.NewResp(nil, 404, "["+request.Method+"] "+request.RequestURI+"访问路由不存在", nil)
	b, err := util.StructToJson(data)
	if err != nil {
		logs.Error(err)
		return
	}
	_, _ = writer.Write(b)
}
