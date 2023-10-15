package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"social-pet/msg"
)

func HttpPostBody(AUrl, ABody, AContenttype string) (string, error) {
	var resp *http.Response
	bytes_req := []byte("")

	bytes_req = []byte(ABody)
	//发送unified order请求.
	req, err := http.NewRequest("POST", AUrl, bytes.NewReader(bytes_req))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", AContenttype)
	//这里的http header的设置是必须设置的.
	req.Header.Set("Content-Type", AContenttype+";charset=utf-8")

	c := http.Client{}
	if resp, err = c.Do(req); err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	//w, _ := resp.Body.Read(buf)

	return string(b), err

}

/*
*
用法
HttpPostStruct("http://baidu.com",struct{ Department string `json:"department"` }{Department: department},&list)
*/
func HttpPostStruct(url string, params interface{}, obj interface{}) (err error) {
	var (
		body string
	)

	if body, err = HttpPostBody(url, StructToJsonString(params), "application/json"); err != nil {
		return
	}
	if err = JsonStrToStruct(body, &obj); err != nil {
		return
	}
	return
}

/*
*
对接his api平台

用法
HisApiPost("http://baidu.com",struct{ Department string `json:"department"` }{Department: department},&list)
*/
func HisApiPost(url string, params interface{}, obj interface{}) (err error) {
	var (
		body       string
		requestStr string
		resp       struct {
			Code    string                   `json:"Code"`
			Message string                   `json:"Message"`
			Records []map[string]interface{} `json:"Records"`
		}
	)
	//如果输入为string，则直接当做参数发送
	if str, ok := params.(string); ok {
		requestStr = str
	} else {
		requestStr = StructToJsonString(params)
	}
	if body, err = HttpPostBody(url, requestStr, "text/plain"); err != nil {
		return
	}
	if err = JsonStrToStruct(body, &resp); err != nil {
		return
	}
	if resp.Code != "0" {
		err = errors.New(resp.Message)
		return
	}
	if obj != nil {
		if err = InterfaceToStruct(resp.Records, obj); err != nil {
			return
		}
	}

	return
}

/*
*
对接beego api平台

用法
YtApiPost("http://baidu.com",struct{ Department string `json:"department"` }{Department: department},&list)
*/
func YtApiPost(url string, params interface{}, obj interface{}, pageResp interface{}) (err error) {
	var (
		body       string
		requestStr string
		resp       msg.Resp
	)
	//如果输入为string，则直接当做参数发送
	if str, ok := params.(string); ok {
		requestStr = str
	} else {
		requestStr = StructToJsonString(params)
	}
	if body, err = HttpPostBody(url, requestStr, "text/plain"); err != nil {
		return
	}
	if err = JsonStrToStruct(body, &resp); err != nil {
		return
	}
	if resp.Code != 0 {
		err = errors.New(resp.Msg)
		return
	}
	if obj != nil {
		if err = InterfaceToStruct(resp.Data, obj); err != nil {
			return
		}
	}
	if pageResp != nil {
		if err = InterfaceToStruct(resp.Data, pageResp); err != nil {
			return
		}
	}
	return
}
