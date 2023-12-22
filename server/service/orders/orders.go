package orders

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/orders"
	ordersReq "github.com/flipped-aurora/gin-vue-admin/server/model/orders/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"gorm.io/gorm"
)

type OrdersService struct {
}

// CreateOrders 创建订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (osService *OrdersService) CreateOrders(os *orders.Orders) (err error) {
	err = global.GVA_DB.Create(os).Error
	return err
}

// DeleteOrders 删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (osService *OrdersService) DeleteOrders(os orders.Orders) (err error) {
	err = global.GVA_DB.Delete(&os).Error
	return err
}

// DeleteOrdersByIds 批量删除订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (osService *OrdersService) DeleteOrdersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]orders.Orders{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrders 更新订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (osService *OrdersService) UpdateOrders(os orders.Orders) (err error) {
	err = global.GVA_DB.Save(&os).Error
	return err
}

// GetOrders 根据id获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
func (osService *OrdersService) GetOrders(id uint) (os orders.Orders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&os).Error
	return
}

// GetOrdersInfoList 分页获取订单记录
// Author [piexlmax](https://github.com/piexlmax)
type OrderList struct {
	Success int64
	Fail    int64
	List    []*orders.Orders
}

func (osService *OrdersService) GetOrdersInfoList(info ordersReq.OrdersSearch) (ol *OrderList, total int64, err error) {
	ol = new(OrderList)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&orders.Orders{})
	var oss []*orders.Orders
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AdminID > 0 {
		var ids []int64
		if global.GVA_DB.Model(&users.Users{}).Where("admin_id = ?", info.AdminID).Pluck("id", &ids).Error == nil {
			db = db.Where("user_id in (?)", ids)
		}
	}
	if info.Account_id != nil {
		db = db.Where("account_id = ?", info.Account_id)
	}
	if info.Order_no != "" {
		db = db.Where("order_no = ?", info.Order_no)
	}
	if info.Direction != nil {
		db = db.Where("direction = ?", info.Direction)
	}
	if info.Volume != nil {
		db = db.Where("volume = ?", info.Volume)
	}
	if info.Price != nil {
		db = db.Where("price = ?", info.Price)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.User_id != nil {
		db = db.Where("user_id = ?", info.User_id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&oss).Error
	for _, al := range oss {
		var u users.Users
		global.GVA_DB.Where("id = ?", *al.User_id).First(&u)
		al.User = u
	}
	ol.List = oss
	db1 := db.Session(&gorm.Session{})
	db2 := db.Session(&gorm.Session{})
	db1.Where("is_win = 1").Count(&ol.Success)
	db2.Where("is_win = 2").Count(&ol.Fail)
	return ol, total, err
}
