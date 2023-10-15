package controllers

import (
	"social-pet/beego/controller"
	"social-pet/global"
	"social-pet/models"
	"social-pet/util"
)

type MessageController struct {
	controller.BaseController
}

func (m *MessageController) Send() {
	var (
		req models.MessageInfo

		err error
	)

	defer func() {
		m.WriteJsonMsgWithError(nil, err)
	}()

	if err = m.ReadJSON(&req); err != nil {
		return
	}

	req.Id = util.NewShortUUId()
	_, err = global.DB.Table(new(models.MessageInfo).TableName()).
		Insert(&req)
}

func (m *MessageController) List() {
	var (
		req struct {
			TelegramId string `json:"telegram_id"`
		}

		err error

		res = make([]models.MessageInfo, 0)
	)
	defer func() {
		m.WriteJsonMsgWithError(res, err)
	}()

	if err = m.ReadJSON(&req); err != nil {
		return
	}

	err = global.DB.Table(new(models.MessageInfo).TableName()).
		Where("telegram_id = ?", req.TelegramId).
		Asc("create_at").
		Find(&res)
}
