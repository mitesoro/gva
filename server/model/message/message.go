package message

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Message struct {
	global.GVA_MODEL
	UserId  int64  `json:"user_id"  form:"user_id" gorm:"column:user_id;comment:用户编号;"` // 用户编号
	Status  int64  `json:"status"  form:"status" gorm:"column:status;comment:是否已读;"`    // 是否已读 1读取 0未读
	Content string `json:"content"  form:"content" gorm:"column:content;comment:内容;"`   // 内容
}
