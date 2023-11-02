// 自动生成模板Recharge
package recharge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 充值 结构体  Recharge
type Recharge struct {
      global.GVA_MODEL
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户编号;"`  //用户编号 
      Amount  *int `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`  //金额 
}


// TableName 充值 Recharge自定义表名 recharge
func (Recharge) TableName() string {
  return "recharge"
}

