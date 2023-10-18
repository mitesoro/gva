package apis

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 短信 结构体  Sms
type ReqSms struct {
	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"` // 手机号
}

// 用户 结构体  Users
type Sms struct {
	global.GVA_MODEL
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"`     // 手机号
	Code     string `json:"code" form:"code" gorm:"column:code;comment:验证码;"`                 // 验证码
	ExpireAt int64  `json:"expire_at" form:"expire_at" gorm:"column:expire_at;comment:过期时间;"` // 过期时间
	Status   int64  `json:"status" form:"status" gorm:"column:status;comment:状态;"`            // 状态 0 可用 1已使用
}
