package util

type Resp struct {
	Code int         `json:"code"`           //状态码
	Msg  string      `json:"msg"`            //返回消息
	Data interface{} `json:"data"`           //数据实体
	Page interface{} `json:"page,omitempty"` //分页数据
}

func NewResp(data interface{}, code int, msg string, page interface{}) (resp *Resp) {
	resp = &Resp{
		Code: code,
		Msg:  msg,
		Data: data,
		Page: page,
	}
	return resp
}

//只传data
func NewSuccessResp(data interface{}) (resp *Resp) {
	resp = &Resp{
		Code: 0,
		Msg:  "",
		Data: data,
	}
	return resp
}

