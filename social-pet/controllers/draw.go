package controllers

import (
	"social-pet/api/baidu"
	"social-pet/beego/controller"
	"social-pet/global"
	"social-pet/models"
	"social-pet/util"
)

/*
 * 步骤
 * 1、选择图片
 * 2、文件上传
 * 3、AI作画
 * 4、铸造NFT
 * 5、完成铸造
 */

type DrawController struct {
	controller.BaseController
}

// AI作画
func (a *DrawController) Draw() {
	var (
		req struct {
			Prompt string `json:"prompt"`
			Image  string `json:"image" valid:"Required"`
		}

		err error

		res struct {
			TaskId string `json:"task_id"`
		}
	)

	defer func() {
		a.WriteJsonMsgWithError(res, err)
	}()

	if err = a.ReadValidJSON(&req); err != nil {
		return
	}

	// //生成AI图片(提交请求)
	if res.TaskId, err = baidu.Draw(req.Prompt, req.Image); err != nil {
		return
	}

	//AI作画记录
	draw := models.DrawInfo{
		Id:     util.NewShortUUId(),
		TaskId: res.TaskId,
		Status: 0,
	}

	if _, err = global.DB.Table(new(models.DrawInfo).TableName()).
		Insert(&draw); err != nil {
		return
	}
}

func (a *DrawController) Get() {
	var (
		req struct {
			TaskId string `json:"task_id"`
		}

		err error

		res models.DrawInfo
	)

	defer func() {
		a.WriteJsonMsgWithError(res, err)
	}()

	if err = a.ReadJSON(&req); err != nil {
		return
	}

	_, err = global.DB.Table(new(models.DrawInfo).TableName()).
		Where("task_id = ?", req.TaskId).
		Get(&res)
}
