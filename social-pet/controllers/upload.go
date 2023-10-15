package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"social-pet/beego/controller"

	"github.com/astaxie/beego/logs"
)

type UploadController struct {
	controller.BaseController
}

// 单文件上传到本地
func (u *UploadController) UploadFile() {
	var (
		res string

		err error
	)

	defer func() {
		u.WriteJsonMsgWithError(res, err)
	}()

	var path string
	path, err = u.SaveUploadFile("file", "./static/img/")

	res = path[1:]
}

// 多文件上传到本地
func (u *UploadController) UploadFiles() {
	var (
		res []string

		err error
	)

	defer func() {
		u.WriteJsonMsgWithError(res, err)
	}()

	var path []string
	path, err = u.SaveUploadFiles("files", "./static/img/")

	for _, p := range path {
		res = append(res, p[1:])
	}
}

type Metadata struct {
	Attributes  []Attributes `json:"attributes"`
	Description string       `json:"description"`
	Image       string       `json:"image"`
	Name        string       `json:"name"`
}

type Attributes struct {
	Value     string `json:"value"`
	TraitType string `json:"trait_type"`
}

type JSONType struct {
	Metadata Metadata `json:"metadata"`
	TokenId  int64    `json:"token_id"`
}

// JSON上传到本地
func (u *UploadController) UploadJson() {
	var (
		req JSONType

		err error

		res string
	)

	defer func() {
		u.WriteJsonMsgWithError(res, err)
	}()

	if err = u.ReadJSON(&req); err != nil {
		return
	}

	// 创建一个文件用于保存
	save := fmt.Sprintf("./static/json/%d.json", req.TokenId)
	var out *os.File
	if out, err = os.Create(save); err != nil {
		logs.Error(err.Error())
	}
	defer out.Close()

	var body []byte
	if body, err = json.Marshal(req.Metadata); err != nil {
		return
	}

	// 然后将响应流和文件流对接起来
	if _, err = io.Copy(out, bytes.NewReader(body)); err != nil {
		logs.Error(err.Error())
	}

	res = save[8:]
}
