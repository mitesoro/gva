// 自动生成模板Orders
package orders

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
)

// 订单 结构体  Orders
type Orders struct {
	global.GVA_MODEL
	User_id        *int            `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户编号;"`                         // 用户编号
	Account_id     *int            `json:"account_id" form:"account_id" gorm:"column:account_id;comment:账户;"`                  // 账户
	Order_no       string          `json:"order_no" form:"order_no" gorm:"column:order_no;comment:订单号;"`                       // 订单号
	Direction      *int            `json:"direction" form:"direction" gorm:"column:direction;comment:类型;"`                     // 类型 1买2卖
	Volume         *int            `json:"volume" form:"volume" gorm:"column:volume;comment:手;size:手;"`                        // 手
	Price          *int            `json:"price" form:"price" gorm:"column:price;comment:价格;"`                                 // 价格
	Bond           *int            `json:"bond" form:"bond" gorm:"column:bond;comment:保证金;"`                                   // 保证金
	Status         *int            `json:"status" form:"status" gorm:"column:status;comment:状态;"`                              // 状态 0下单中 1成功 2取消 3失败 4盈亏 5平仓
	Content        string          `json:"content" form:"content" gorm:"column:content;comment:内容备注;"`                         //  期货下单返回数据
	Amount         *int            `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                              // 盈利金额
	CompleteAt     model.LocalTime `json:"complete_at" form:"complete_at" gorm:"column:complete_at;comment:平仓时间;"`             // 平仓时间
	SymbolID       string          `json:"symbol_id" form:"symbol_id" gorm:"column:symbol_id;comment:合约编号;"`                   // 合约品种编号
	SymbolName     string          `json:"symbol_name" form:"symbol_name" gorm:"column:symbol_name;comment:合约名称;"`             // 合约品种名称
	OrderRef       int             `json:"order_ref" form:"order_ref" gorm:"column:order_ref;comment:合约单;"`                    // 合约单
	OrderSysID     int             `json:"order_sys_id" form:"order_sys_id" gorm:"column:order_sys_id;comment:下单编号;"`          // 下单编号
	SuccessAt      model.LocalTime `json:"success_at" form:"success_at" gorm:"column:success_at;comment:成交时间;"`                // 成交时间
	SuccessPrice   int64           `json:"success_price" form:"success_price" gorm:"column:success_price;comment:成交价;"`        // 成交价
	DecrAmount     int64           `json:"decr_amount" form:"decr_amount" gorm:"column:decr_amount;comment:扣除金钱;"`             // 扣除金钱
	Fee            int64           `json:"fee" form:"fee" gorm:"column:fee;comment:手续费;"`                                      // 手续费
	IsWin          int64           `json:"is_win" form:"is_win" gorm:"column:is_win;comment:平仓盈亏;"`                            // 平仓盈亏 1赢2亏
	WinAmount      int64           `json:"win_amount" form:"win_amount" gorm:"column:win_amount;comment:平仓盈亏金额;"`              // 平仓盈亏金额
	ThirdDirection int             `json:"third_direction" form:"third_direction" gorm:"column:third_direction;comment:外部类型;"` // 期货类型 1买2卖
	ThirdPrice     int             `json:"third_price" form:"third_price" gorm:"column:third_price;comment:外部价格;"`             // 期货价格
	ClosePrice     int             `json:"close_price" form:"close_price" gorm:"column:close_price;comment:平仓价格;"`             // 平仓价格
	User           users.Users
}

// TableName 订单 Orders自定义表名 orders
func (Orders) TableName() string {
	return "orders"
}
