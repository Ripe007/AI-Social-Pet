package models

import "social-pet/beego/bmodel"

// AI作画表
type DrawInfo struct {
	Id                string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt          bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt         bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt         bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	SourceImageBase64 string           `json:"source_image_base64" xorm:"TEXT"`    //原始图片源数据
	TaskId            string           `json:"task_id" xorm:"VARCHAR(30)"`         //生成图片，提交请求ID
	ImageUrl          string           `json:"image_url" xorm:"VARCHAR(200)"`      //生成图片，获取图片URL
	AnimalIdentify    string           `json:"animal_identify" xorm:"VARCHAR(50)"` //物种名称
	Status            int64            `json:"status" xorm:"int"`                  //状态
}

func (a *DrawInfo) TableName() string {
	return "draw_info"
}

// 群组
type TelegramInfo struct {
	Id             string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt       bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt      bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt      bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	AnimalId       string           `json:"animal_id" xorm:"VARCHAR(100)"`   //物种ID
	NickName       string           `json:"nick_name" xorm:"VARCHAR(50)"`    //昵称
	Description    string           `json:"description" xorm:"VARCHAR(200)"` //简述
	TotalUserCount int64            `json:"total_user_count" xorm:"int"`     //总人数
}

func (t *TelegramInfo) TableName() string {
	return "telegram_info"
}

// 用户-群组关系表
type UserTelegramInfo struct {
	Id         string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt   bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt  bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt  bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	UserAddr   string           `json:"user_addr" xorm:"VARCHAR(50)"`    //用户地址
	TelegramId string           `json:"telegram_id" xorm:"VARCHAR(100)"` //群组ID
}

func (u *UserTelegramInfo) TableName() string {
	return "user_telegram_info"
}

// 物种
type AnimalInfo struct {
	Id        string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt  bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	Name      string           `json:"name" xorm:"VARCHAR(50)"` //物种名称
}

func (a *AnimalInfo) TableName() string {
	return "animal_info"
}

// 铸造数据
type MintInfo struct {
	Id        string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt  bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	UserAddr  string           `json:"user_addr" xorm:"VARCHAR(50)"` //用户地址
	TxId      string           `json:"tx_id" xorm:"VARCHAR(100)"`    //交易哈希
}

func (a *MintInfo) TableName() string {
	return "mint_info"
}

// 动态社区内容
type CommunityContextInfo struct {
	Id          string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt    bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt   bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt   bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	UserAddr    string           `json:"user_addr" xorm:"VARCHAR(50)"`         //发布用户地址
	Context     string           `json:"context" xorm:"TEXT"`                  //文本内容
	ImgUrls     []string         `json:"img_urls" xorm:"TEXT"`                 //图片
	Status      int64            `json:"status" xorm:"int not null default 0"` //1:自发 2:转发 -1:已删除
	ForkContext string           `json:"fork_context" xorm:"TEXT"`             //转发文本内容
	SourceId    string           `json:"source_id" xorm:"VARCHAR(100)"`        //转发自源文件
}

func (c *CommunityContextInfo) TableName() string {
	return "community_context_info"
}

// 社区动态配置数据记录
type CommunityConfigInfo struct {
	Id           string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt     bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt    bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt    bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	SourceId     string           `json:"source_id" xorm:"VARCHAR(100)"`               //源动态ID
	LikeCount    int64            `json:"like_count" xorm:"int not null default 0"`    //点赞次数
	CommentCount int64            `json:"comment_count" xorm:"int not null default 0"` //评论次数
	ForwardCount int64            `json:"forward_count" xorm:"int not null default 0"` //转发次数
}

func (c *CommunityConfigInfo) TableName() string {
	return "community_config_info"
}

// 社区动态点赞表
type CommunityLikeInfo struct {
	Id        string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt  bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	SourceId  string           `json:"source_id" xorm:"VARCHAR(100)"`        //源动态ID
	UserAddr  string           `json:"user_addr" xorm:"VARCHAR(50)"`         //用户地址
	Status    int64            `json:"status" xorm:"int not null default 0"` //1:确定点赞 -1:取消点赞
}

func (c *CommunityLikeInfo) TableName() string {
	return "community_like_info"
}

// 社区动态评论表
type CommunityCommentInfo struct {
	Id            string                  `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt      bmodel.LocalTime        `xorm:"created" json:"create_at"`
	UpdatedAt     bmodel.LocalTime        `xorm:"updated" json:"updated_at"`
	DeletedAt     bmodel.LocalTime        `xorm:"deleted" json:"deleted_at"`
	SourceId      string                  `json:"source_id" xorm:"VARCHAR(100)"`       //源内容ID
	LastCommentId string                  `json:"last_comment_id" xorm:"VARCHAR(100)"` //上一条评论ID
	UserAddr      string                  `json:"user_addr" xorm:"VARCHAR(50)"`        //用户地址
	Context       string                  `json:"context" xorm:"TEXT"`                 //文本内容
	Child         []*CommunityCommentInfo `xorm:"-"`
}

func (c *CommunityCommentInfo) TableName() string {
	return "community_comment_info"
}

// 聊天消息表
type MessageInfo struct {
	Id         string           `json:"id" xorm:"not null pk unique VARCHAR(100) 'id'" `
	CreateAt   bmodel.LocalTime `xorm:"created" json:"create_at"`
	UpdatedAt  bmodel.LocalTime `xorm:"updated" json:"updated_at"`
	DeletedAt  bmodel.LocalTime `xorm:"deleted" json:"deleted_at"`
	UserAddr   string           `json:"user_addr" xorm:"VARCHAR(50)"`    //用户地址
	Message    string           `json:"message" xorm:"TEXT"`             //消息内容
	TelegramId string           `json:"telegram_id" xorm:"VARCHAR(100)"` //群组ID
}

func (m *MessageInfo) TableName() string {
	return "message_info"
}
