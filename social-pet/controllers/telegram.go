package controllers

import (
	"social-pet/beego/controller"
	"social-pet/global"
	"social-pet/models"
)

type TelegramController struct {
	controller.BaseController
}

func (t *TelegramController) List() {
	var (
		req struct {
			UserAddr string `json:"user_addr"`
		}
		err error

		res = make([]models.TelegramInfo, 0)
	)
	defer func() {
		t.WriteJsonMsgWithError(res, err)
	}()

	if err = t.ReadJSON(&req); err != nil {
		return
	}

	err = global.DB.Table(new(models.TelegramInfo).TableName()).
		Join("INNER", new(models.UserTelegramInfo).TableName(), "telegram_info.id=user_telegram_info.telegram_id").
		Where("user_telegram_info.user_addr = ?", req.UserAddr).
		Cols("telegram_info.*").
		Desc("telegram_info.create_at").
		Find(&res)
}

func (t *TelegramController) UserList() {
	var (
		req struct {
			TelegramId string `json:"telegram_id"`
		}
		err error

		res = make([]models.UserTelegramInfo, 0)
	)
	defer func() {
		t.WriteJsonMsgWithError(res, err)
	}()

	if err = t.ReadJSON(&req); err != nil {
		return
	}

	err = global.DB.Table(new(models.UserTelegramInfo).TableName()).
		Where("telegram_id = ?", req.TelegramId).
		Desc("create_at").
		Find(&res)
}
