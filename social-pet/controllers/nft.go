package controllers

import (
	"errors"
	"social-pet/beego/controller"
	"social-pet/global"
	"social-pet/models"
	"social-pet/util"
)

type NFTController struct {
	controller.BaseController
}

func (a *NFTController) Count() {
	var (
		req struct {
			UserAddr string `json:"user_addr"`
		}

		err error

		res int64
	)

	defer func() {
		a.WriteJsonMsgWithError(res, err)
	}()

	if err = a.ReadJSON(&req); err != nil {
		return
	}

	res, err = global.DB.Table(new(models.MintInfo).TableName()).
		Where("user_addr = ?", req.UserAddr).
		Count(new(models.MintInfo))

}

func (a *NFTController) Mint() {
	var (
		req struct {
			TxId     string `json:"tx_id" valid:"Required"`
			UserAddr string `json:"user_addr" valid:"Required"`
			Identify string `json:"identify" valid:"Required"`
		}

		sess = global.DB.NewSession()

		err error
	)

	defer func() {
		a.WriteJsonMsgWithError(nil, err)
	}()
	defer sess.Close()

	if err = a.ReadValidJSON(&req); err != nil {
		return
	}

	sess.Begin()

	//判断是否已铸造
	var mint models.MintInfo
	if _, err = sess.Table(new(models.MintInfo).TableName()).
		Where("tx_id = ?", req.TxId).
		Get(&mint); err != nil {
		return
	}

	if mint.Id != "" {
		err = errors.New("tx is exist")
		return
	} else {
		//生成铸造数据
		mint = models.MintInfo{
			Id:       util.NewShortUUId(),
			UserAddr: req.UserAddr,
			TxId:     req.TxId,
		}

		if _, err = sess.Table(new(models.MintInfo).TableName()).
			Insert(&mint); err != nil {
			sess.Rollback()
			return
		}
	}

	//判断物种是否存在，不存在则新建物种
	var animal models.AnimalInfo
	if _, err = sess.Table(new(models.AnimalInfo).TableName()).
		Where("name = ?", req.Identify).
		Get(&animal); err != nil {
		return
	}

	if animal.Id == "" {
		//生成物种信息
		animal = models.AnimalInfo{
			Id:   util.NewShortUUId(),
			Name: req.Identify,
		}

		if _, err = sess.Table(new(models.AnimalInfo).TableName()).
			Insert(&animal); err != nil {
			sess.Rollback()
			return
		}
	}

	//判断群组存不存在,不存在则新建群组
	var telegram models.TelegramInfo
	if _, err = sess.Table(new(models.TelegramInfo).TableName()).
		Where("animal_id = ?", animal.Id).
		Get(&telegram); err != nil {
		return
	}

	if telegram.Id == "" {
		//生成群组信息
		telegram = models.TelegramInfo{
			Id:             util.NewShortUUId(),
			AnimalId:       animal.Id,
			NickName:       req.Identify,
			Description:    "",
			TotalUserCount: 1,
		}

		if _, err = sess.Table(new(models.TelegramInfo).TableName()).
			Insert(&telegram); err != nil {
			sess.Rollback()
			return
		}
	} else {
		//更新群组总人数
		telegram.TotalUserCount += 1
		if _, err = sess.Table(new(models.TelegramInfo).TableName()).
			ID(telegram.Id).
			Update(&telegram); err != nil {
			sess.Rollback()
			return
		}
	}

	//用户加入群组
	user_telegram := models.UserTelegramInfo{
		Id:         util.NewShortUUId(),
		UserAddr:   req.UserAddr,
		TelegramId: telegram.Id,
	}

	if _, err = sess.Table(new(models.UserTelegramInfo).TableName()).
		Insert(&user_telegram); err != nil {
		sess.Rollback()
		return
	}

	sess.Commit()
}
