// 自动生成模板Orders
package orders

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 订单 结构体  Orders
type Orders struct {
      global.GVA_MODEL
      User_id  *int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户编号;"`  //用户编号 
      Account_id  *int `json:"account_id" form:"account_id" gorm:"column:account_id;comment:账户;"`  //账户 
      Order_no  string `json:"order_no" form:"order_no" gorm:"column:order_no;comment:订单号;"`  //订单号 
      Direction  *int `json:"direction" form:"direction" gorm:"column:direction;comment:类型;"`  //类型 
      Volume  *int `json:"volume" form:"volume" gorm:"column:volume;comment:手;size:手;"`  //手 
      Price  *int `json:"price" form:"price" gorm:"column:price;comment:价格;"`  //价格 
}


// TableName 订单 Orders自定义表名 orders
func (Orders) TableName() string {
  return "orders"
}

