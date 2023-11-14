// 自动生成模板Symbol
package symbols

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 合约品种 结构体  Symbol
type Symbol struct {
	global.GVA_MODEL
	Name              string `json:"name" form:"name" gorm:"column:name;comment:名称;"`                                                  // 名称
	Code              string `json:"code" form:"code" gorm:"column:code;comment:代码;"`                                                  // 代码
	Multiple          *int   `json:"multiple" form:"multiple" gorm:"column:multiple;comment:倍数;"`                                      // 倍数
	Bond              *int   `json:"bond" form:"bond" gorm:"column:bond;comment:保证金;"`                                                 // 保证金
	Type              int    `json:"type" form:"type" gorm:"column:type;comment:类型;"`                                                  // 类型 0 固定值 1 百分比
	Amount            int    `json:"amount" form:"amount" gorm:"column:amount;comment:固定值;"`                                           // 固定值
	Status            *bool  `json:"status" form:"status" gorm:"column:status;comment:状态;"`                                            // 状态
	PointSuccess      *int   `json:"point_success" form:"point_success" gorm:"column:point_success;comment:止赢点位;"`                     // 止赢点位
	PointSuccessPrice *int   `json:"point_success_price" form:"point_success_price" gorm:"column:point_success_price;comment:止赢点位价格;"` // 止赢点位价格
	PointFail         *int   `json:"point_fail" form:"point_fail" gorm:"column:point_fail;comment:止损点位;"`                              // 止损点位
	PointFailPrice    *int   `json:"point_fail_price" form:"point_fail_price" gorm:"column:point_fail_price;comment:止损点位价格;"`          // 止损点位价格
	Days              string `json:"days" form:"days" gorm:"column:days;comment:开盘也是时间;"`                                              // 开盘特殊时间
	Times             string `json:"times" form:"times" gorm:"column:times;comment:时间限制;"`                                             // 时间限制
}

// TableName 合约品种 Symbol自定义表名 symbol
func (Symbol) TableName() string {
	return "symbol"
}
