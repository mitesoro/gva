package apis

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apis"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	usersReq "github.com/flipped-aurora/gin-vue-admin/server/model/users/request"
	"time"
)

type ApisService struct {
}

// CreateSms 创建短信记录
func (uService *ApisService) CreateSms(u *apis.Sms) (err error) {
	err = global.GVA_DB.Create(u).Error
	return err
}

// CheckSms 验证sms
func (uService *ApisService) CheckSms(phone, code string) bool {
	var sms apis.Sms
	if err := global.GVA_DB.Where("phone = ? and code = ?", phone, code).Order("id DESC").
		First(&sms).Error; err == nil {
		if sms.ExpireAt < time.Now().Unix() && sms.Status == 0 {
			sms.Status = 1
			global.GVA_DB.Save(&sms)
			return true
		}
	}
	return false
}

// GetUsersInfoList 分页获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (uService *ApisService) GetUsersInfoList(info usersReq.UsersSearch) (list []users.Users, total int64, err error) {
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
