package controllers

import (
	"errors"
	"social-pet/beego/controller"
	"social-pet/global"
	"social-pet/handle"
	"social-pet/models"
	"social-pet/util"
)

type CommunityController struct {
	controller.BaseController
}

// 发布/转发社区动态
func (c *CommunityController) Publish() {
	var (
		req models.CommunityContextInfo

		sess = global.DB.NewSession()

		err error
	)

	defer func() {
		c.WriteJsonMsgWithError(nil, err)
	}()

	defer sess.Close()

	if err = c.ReadJSON(&req); err != nil {
		return
	}

	//判断用户是否持有NFT
	// var balance *big.Int
	// if balance, err = example.BalanceOf(req.UserAddr); err != nil {
	// 	return
	// }

	/*
	 如果 x < 0 返回 -1，如果 x == 0 返回 0，如果 x > 0 返回 +1
	*/
	// if balance.Sign() <= 0 {
	// 	err = errors.New("no nft")
	// 	return
	// }

	sess.Begin()

	req.Id = util.NewShortUUId()

	//判断该文本是转发or自发
	var config models.CommunityConfigInfo
	if req.SourceId == "" {
		//自发
		req.Status = 1

		req.SourceId = req.Id

		//构造配置
		config = models.CommunityConfigInfo{
			Id:       util.NewShortUUId(),
			SourceId: req.Id,
		}

		//新增配置
		if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
			Insert(&config); err != nil {
			sess.Rollback()
			return
		}
	} else {
		//转发
		req.Status = 2

		var source models.CommunityContextInfo
		if _, err = sess.Table(new(models.CommunityContextInfo).TableName()).
			ID(req.SourceId).
			Get(&source); err != nil {
			return
		}

		if source.Id == "" {
			err = errors.New("source not found")
			return
		}

		//判断该文本是否是自己发布，自己发布不能转发
		if source.UserAddr == req.UserAddr {
			err = errors.New("Insufficient permissions")
			return
		}

		//查询配置
		if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
			Where("source_id = ?", req.SourceId).
			Get(&config); err != nil {
			return
		}

		//更新配置数据
		config.ForwardCount += 1
		if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
			ID(config.Id).
			Cols("forward_count").
			Update(&config); err != nil {
			sess.Rollback()
			return
		}
	}

	if _, err = sess.Table(new(models.CommunityContextInfo).TableName()).
		Insert(&req); err != nil {
		sess.Rollback()
		return
	}

	sess.Commit()
}

// 点赞/取消点赞
func (c *CommunityController) Like() {
	var (
		req models.CommunityLikeInfo

		sess = global.DB.NewSession()

		err error
	)

	defer func() {
		c.WriteJsonMsgWithError(nil, err)
	}()

	defer sess.Close()

	if err = c.ReadJSON(&req); err != nil {
		return
	}

	sess.Begin()

	var source models.CommunityContextInfo
	if _, err = sess.Table(new(models.CommunityContextInfo).TableName()).
		ID(req.SourceId).
		Get(&source); err != nil {
		return
	}

	if source.Id == "" {
		err = errors.New("source not found")
		return
	}

	//判断该文本是否是自己发布，自己发布不能点赞
	if source.UserAddr == req.UserAddr {
		err = errors.New("Insufficient permissions")
		return
	}

	//获取动态配置数据
	var config models.CommunityConfigInfo
	if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
		Where("source_id = ?", req.SourceId).
		Get(&config); err != nil {
		return
	}

	//判断用户是否已点赞
	var like models.CommunityLikeInfo
	if _, err = sess.Table(new(models.CommunityLikeInfo).TableName()).
		Where("source_id = ?", req.SourceId).
		Get(&like); err != nil {
		return
	}

	if like.Id == "" {
		config.LikeCount += 1

		//点赞
		req.Id = util.NewShortUUId()
		req.Status = 1
		if _, err = sess.Table(new(models.CommunityLikeInfo).TableName()).
			Insert(&req); err != nil {
			sess.Rollback()
			return
		}
	} else {
		if like.Status == 1 {
			//若上一步是点赞状态，则当前置为取消点赞状态

			//取消点赞
			config.LikeCount -= 1

			like.Status = -1
		} else if like.Status == -1 {
			//若上一步是取消点赞状态，则当前置为点赞状态

			//点赞
			config.LikeCount += 1

			like.Status = 1
		}

		if _, err = sess.Table(new(models.CommunityLikeInfo).TableName()).
			ID(like.Id).
			Cols("status").
			Update(&like); err != nil {
			sess.Rollback()
			return
		}
	}

	if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
		ID(config.Id).
		Cols("like_count").
		Update(&config); err != nil {
		sess.Rollback()
	}

	sess.Commit()
}

// 添加评论
func (c *CommunityController) CommentAdd() {
	var (
		req models.CommunityCommentInfo

		sess = global.DB.NewSession()

		err error
	)

	defer func() {
		c.WriteJsonMsgWithError(nil, err)
	}()

	defer sess.Close()

	if err = c.ReadJSON(&req); err != nil {
		return
	}

	sess.Begin()

	var source models.CommunityContextInfo
	if _, err = sess.Table(new(models.CommunityContextInfo).TableName()).
		ID(req.SourceId).
		Get(&source); err != nil {
		return
	}

	if source.Id == "" {
		err = errors.New("source not found")
		return
	}

	req.Id = util.NewShortUUId()
	if _, err = sess.Table(new(models.CommunityCommentInfo).TableName()).
		Insert(&req); err != nil {
		sess.Rollback()
		return
	}

	var config models.CommunityConfigInfo
	if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
		Where("source_id = ?", req.SourceId).
		Get(&config); err != nil {
		return
	}

	//更新配置数据
	config.CommentCount += 1
	if _, err = sess.Table(new(models.CommunityConfigInfo).TableName()).
		ID(config.Id).
		Cols("comment_count").
		Update(&config); err != nil {
		sess.Rollback()
		return
	}

	sess.Commit()

}

// 评论列表
func (c *CommunityController) CommentList() {
	var (
		req struct {
			SourceId string `json:"source_id"`
		}

		err error

		list = make([]*models.CommunityCommentInfo, 0)
	)

	defer func() {
		c.WriteJsonMsgWithError(list, err)
	}()

	if err = c.ReadJSON(&req); err != nil {
		return
	}

	list, err = handle.GetCommunityAllSubordinates(req.SourceId)
}

// 社区动态
func (c *CommunityController) List() {
	var (
		req struct {
			SourceId string `json:"source_id"`
		}

		err error

		list = make([]struct {
			Context models.CommunityContextInfo `json:"context" xorm:"extends"`
			Config  models.CommunityConfigInfo  `json:"config" xorm:"extends"`
		}, 0)
	)

	defer func() {
		c.WriteJsonMsgWithError(list, err)
	}()

	if err = c.ReadJSON(&req); err != nil {
		return
	}

	coll := global.DB.Table(new(models.CommunityContextInfo).TableName()).
	Join("INNER", new(models.CommunityConfigInfo).TableName(), "community_context_info.id=community_config_info.source_id")

	if req.SourceId != "" {
		coll.Where("community_context_info.id = ?", req.SourceId)
	}

	err = coll.Desc("community_context_info.create_at").
		Find(&list)
}

func (c *CommunityController) LikeGet() {
	var (
		req struct {
			UserAddr string `json:"user_addr"`
			SourceId string `json:"source_id"`
		}

		err error

		res models.CommunityLikeInfo
	)
	defer func() {
		c.WriteJsonMsgWithError(res, err)
	}()

	if err = c.ReadJSON(&req); err != nil {
		return
	}

	_, err = global.DB.Table(new(models.CommunityLikeInfo).TableName()).
		Where("user_addr = ?", req.UserAddr).
		Where("source_id = ?", req.SourceId).
		Get(&res)

}
