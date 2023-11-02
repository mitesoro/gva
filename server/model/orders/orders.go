// 自动生成模板Orders
package orders

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 订单 结构体  Orders
type Orders struct {
	global.GVA_MODEL
	User_id    *int   `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户编号;"`        // 用户编号
	Account_id *int   `json:"account_id" form:"account_id" gorm:"column:account_id;comment:账户;"` // 账户
	Order_no   string `json:"order_no" form:"order_no" gorm:"column:order_no;comment:订单号;"`      // 订单号
	Direction  *int   `json:"direction" form:"direction" gorm:"column:direction;comment:类型;"`    // 类型
	Volume     *int   `json:"volume" form:"volume" gorm:"column:volume;comment:手;size:手;"`       // 手
	Price      *int   `json:"price" form:"price" gorm:"column:price;comment:价格;"`                // 价格
	Bond       *int   `json:"bond" form:"bond" gorm:"column:bond;comment:保证金;"`                  // 保证金
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;"`             // 状态 0下单中 1成功 2取消 3失败
	Content    string `json:"content" form:"content" gorm:"column:content;comment:内容备注;"`        //  期货下单返回数据
	Amount     *int   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`             // 盈利金额
}

// TableName 订单 Orders自定义表名 orders
func (Orders) TableName() string {
	return "orders"
}
