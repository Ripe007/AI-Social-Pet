package handle

import (
	"reflect"
	"social-pet/global"
	"social-pet/models"
)

func GetCommunityAllSubordinates(communityId string) (res []*models.CommunityCommentInfo, err error) {
	//定义指针切片用来存储所有用户
	var comments []*models.CommunityCommentInfo

	//查询所有一级用户
	if err = global.DB.Table(new(models.CommunityCommentInfo)).
		Where("source_id = ?", communityId).
		Where("last_comment_id = ''").
		Find(&comments); err != nil {
		return
	}

	//判断是否存在数据,存在进行树状图重构
	if reflect.ValueOf(comments).IsValid() {
		//将一级用户传递给回调函数
		if res, err = tree(comments); err != nil {
			return
		}
	}
	return
}

// 生成树结构
func tree(comments []*models.CommunityCommentInfo) ([]*models.CommunityCommentInfo, error) {
	var err error

	if len(comments) == 0 || comments == nil {
		return nil, nil
	}
	if reflect.ValueOf(comments).IsValid() {
		//循环所有一级用户
		for k, _ := range comments {
			//定义子节点目录
			var nodes []*models.CommunityCommentInfo
			//查询所有该用户下的所有子用户
			if err = global.DB.Table(new(models.CommunityCommentInfo).TableName()).
				Where("last_comment_id = ?", comments[k].Id).
				Find(&nodes); err != nil {
				return nil, err
			}

			//将子用户的数据循环赋值给父用户
			for kk, _ := range nodes {
				comments[k].Child = append(comments[k].Child, nodes[kk])
			}
			//将刚刚查询出来的子用户进行递归,查询出三级用户和四级用户
			if len(comments) > 0 {
				tree(nodes)
			}
		}
	}
	return comments, nil
}
