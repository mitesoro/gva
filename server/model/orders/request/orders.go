package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orders"
	"time"
)

type OrdersSearch struct {
	orders.Orders
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	AdminID        int        `json:"admin_id" form:"admin_id" gorm:"column:admin_id;comment:上级;"` // 上级编号
	request.PageInfo
}
