// 自动生成模板Users
package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户 结构体  Users
type Users struct {
	global.GVA_MODEL
	Phone           string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:129;"`                          // 手机号
	Password        string `json:"password" form:"password" gorm:"column:password;comment:密码;"`                           // 密码
	Nickname        string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;"`                           // 昵称
	Avatar          string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;"`                                 // 头像
	OrderType       int    `json:"order_type" form:"order_type" gorm:"column:order_type;comment:订单类型;"`                   // 下单正反手
	Volume          int    `json:"volume" form:"volume" gorm:"column:volume;comment:手;size:手;"`                           // 手
	Rate            int    `json:"rate" form:"rate" gorm:"column:rate;comment:盈亏比;"`                                      // 盈亏比
	Amount          int    `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                 // 总金额
	AvailableAmount int    `json:"available_amount" form:"available_amount" gorm:"column:available_amount;comment:可用金额;"` // 可用金额
	FreezeAmount    int    `json:"freeze_amount" form:"freeze_amount" gorm:"column:freeze_amount;comment:冻结金额;"`          // 冻结金额
}

// TableName 用户 Users自定义表名 users
func (Users) TableName() string {
	return "users"
}
