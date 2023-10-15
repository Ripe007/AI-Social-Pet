package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"mime/multipart"
	"path"
	"social-pet/beego/validation"

	"social-pet/errors"
	"social-pet/msg"
	"social-pet/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

// 跨包继承,面向接口
type BaseControllerInterface interface {
	beego.ControllerInterface
	//获取请求ip
	GetIP() string
	////获取用户头部token
	//GetUserToken() (token string, err error)
	//读取请求参数json
	ReadJSON(params interface{}) (err error)
	//读取请求参数json并校验
	ReadValidJSON(params interface{}) (err error)
	//读取请求参数xml
	ReadXML(params interface{}) (err error)
	//响应客户端
	WriteJsonMsgWithError(data interface{}, err error)
	//带分页的数据
	WritePageJsonMsgWithError(data interface{}, page interface{}, err error)
	//保存上传文件
	SaveUploadFile(fileKey, savePath string) (filePath string, err error)
}

// 在beego.Controller上的进一步封装，主要进行一些验证
type BaseController struct {
	beego.Controller
	responseData  interface{} //回写
	responseError error       //接口报错
}

func (c *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
}

// 获取ip
func (c *BaseController) GetIP() string {
	return util.RealIP(c.Ctx.Request)
}

// 从参数中读取数据
func (c *BaseController) ReadJSON(params interface{}) (err error) {
	if len(c.Ctx.Input.RequestBody) == 0 {
		return
	}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		err = errors.ParamInvalid.Wrap(err, "参数解析错误")
		return
	}
	return
}

// 从参数中读取数据返回map对像
func (c *BaseController) ReadJSONMap() (retmap map[string]interface{}, err error) {
	if len(c.Ctx.Input.RequestBody) == 0 {
		return
	}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &retmap); err != nil {
		err = errors.ParamInvalid.Wrap(err, "参数解析错误")
		return
	}
	return
}

// 从参数中读取数据并校验
func (c *BaseController) ReadValidJSON(params interface{}) (err error) {
	var (
		isValid bool   //校验是否通过
		errMsg  string //错误信息
	)
	if err = c.ReadJSON(params); err != nil {
		return
	}
	valid := validation.Validation{}
	if isValid, err = valid.Valid(params); err != nil {
		err = errors.ParamInvalid.Wrap(err, "校验参数错误")
		return
	}
	if !isValid {
		for _, e := range valid.Errors {
			errMsg += fmt.Sprintf("%s%s;", e.Field, e.Message)
		}
		err = errors.ParamInvalid.New(errMsg)
	}
	return
}

// 从参数中读取数据
func (c *BaseController) ReadXML(params interface{}) (err error) {
	if len(c.Ctx.Input.RequestBody) == 0 {
		return
	}
	if err = xml.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		err = errors.ParamInvalid.Wrap(err, "参数解析错误")
		return
	}
	return
}

// 返回json数据
func (c *BaseController) WriteJsonMsgWithError(data interface{}, err error) {
	var (
		msgStr string //错误信息
		code   int    //错误码
	)
	if err != nil {
		c.responseError = err
		logs.Error(err)
		msgStr = err.Error()
		ErrType := errors.NoType
		if ErrType = errors.GetType(err); ErrType == errors.NoType {
			ErrType = errors.SystemError
		}
		code = int(ErrType)
	}
	resp := msg.NewResp(data, code, msgStr, nil)
	if resp.Data == nil {
		resp.Data = struct{}{}
	}
	c.responseData = resp
	c.Data["json"] = resp
	c.ServeJSON()
}

// 返回json数据
func (c *BaseController) WritePageJsonMsgWithError(data interface{}, page interface{}, err error) {
	var (
		msgStr string //错误信息
		code   int    //错误码
	)
	if err != nil {
		c.responseError = err
		logs.Error(err)
		msgStr = err.Error()
		ErrType := errors.NoType
		if ErrType = errors.GetType(err); ErrType == errors.NoType {
			ErrType = errors.SystemError
		}
		code = int(ErrType)
	}
	resp := msg.NewResp(data, code, msgStr, page)
	if resp.Data == nil {
		resp.Data = struct{}{}
	}
	c.responseData = resp
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *BaseController) SaveUploadFile(fileKey, savePath string) (filePath string, err error) {
	var (
		file       multipart.File
		fileHeader *multipart.FileHeader
	)
	if file, fileHeader, err = c.GetFile(fileKey); err != nil { //获取上传的文件
		err = errors.SystemError.Wrap(err, "获取文件失败")
		return
	}
	defer file.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	ext := path.Ext(fileHeader.Filename)
	//构造文件名称
	fileName := fmt.Sprintf("%s%s", util.NewShortUUId(), ext)
	filePath = savePath + fileName
	if err = c.SaveToFile("file", filePath); err != nil {
		err = errors.SystemError.Wrap(err, "保存文件失败")
		return
	}
	return
}

func (c *BaseController) SaveUploadFiles(fileKey, savePath string) (filePath []string, err error) {
	files := make([]*multipart.FileHeader, 0)
	if files, err = c.GetFiles(fileKey); err != nil { //获取上传的文件
		err = errors.SystemError.Wrap(err, "获取文件失败")
		return
	}

	for i, _ := range files {
		var file multipart.File
		//for each fileheader, get a handle to the actual file
		if file, err = files[i].Open(); err != nil {
			err = errors.SystemError.Wrap(err, "打开文件失败")
			return
		}
		defer file.Close()

		ext := path.Ext(files[i].Filename)
		//构造文件名称
		fileName := fmt.Sprintf("%s%s", util.NewShortUUId(), ext)
		path := savePath + fileName
		if err = c.SaveToFile("files", path); err != nil {
			err = errors.SystemError.Wrap(err, "保存文件失败")
			return

		}

		filePath = append(filePath, path)
	}

	return
}
