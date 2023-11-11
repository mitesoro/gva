// 自动生成模板Alog
package alog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
)

// 金额纪录 结构体  Alog
type Alog struct {
	global.GVA_MODEL
	User_id     *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户编号"`              //用户编号
	Amount_type *int `json:"amount_type" form:"amount_type" gorm:"column:amount_type;comment:金额类型;"` //金额类型
	Amount      *int `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                  //金额
	Cur_amount  *int `json:"cur_amount" form:"cur_amount" gorm:"column:cur_amount;comment:当前金额;"`    //当前金额
	User        users.Users
}

// TableName 金额纪录 Alog自定义表名 amount_log
func (Alog) TableName() string {
	return "amount_log"
}
