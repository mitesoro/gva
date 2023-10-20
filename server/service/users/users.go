package users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	usersReq "github.com/flipped-aurora/gin-vue-admin/server/model/users/request"
)

type UsersService struct {
}

// CreateUsers 创建用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UsersService) CreateUsers(u *users.Users) (err error) {
	err = global.GVA_DB.Create(u).Error
	return err
}

// DeleteUsers 删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UsersService) DeleteUsers(u users.Users) (err error) {
	err = global.GVA_DB.Delete(&u).Error
	return err
}

// DeleteUsersByIds 批量删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UsersService) DeleteUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]users.Users{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUsers 更新用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UsersService) UpdateUsers(u users.Users) (err error) {
	err = global.GVA_DB.Save(&u).Error
	return err
}

// GetUsers 根据id获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UsersService) GetUsers(id uint) (u users.Users, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&u).Error
	return
}

// GetUsersByPhone 根据手机号获取用户记录
func (uService *UsersService) GetUsersByPhone(phone string) (u users.Users, err error) {
	err = global.GVA_DB.Where("phone = ?", phone).First(&u).Error
	return
}

// GetUsersInfoList 分页获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *UsersService) GetUsersInfoList(info usersReq.UsersSearch) (list []users.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&users.Users{})
	var us []users.Users
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["phone"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&us).Error
	return us, total, err
}
