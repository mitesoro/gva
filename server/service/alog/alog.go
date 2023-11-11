package alog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alog"
	alogReq "github.com/flipped-aurora/gin-vue-admin/server/model/alog/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
)

type AlogService struct {
}

// CreateAlog 创建金额纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (alService *AlogService) CreateAlog(al *alog.Alog) (err error) {
	err = global.GVA_DB.Create(al).Error
	return err
}

// DeleteAlog 删除金额纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (alService *AlogService) DeleteAlog(al alog.Alog) (err error) {
	err = global.GVA_DB.Delete(&al).Error
	return err
}

// DeleteAlogByIds 批量删除金额纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (alService *AlogService) DeleteAlogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]alog.Alog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAlog 更新金额纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (alService *AlogService) UpdateAlog(al alog.Alog) (err error) {
	err = global.GVA_DB.Save(&al).Error
	return err
}

// GetAlog 根据id获取金额纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (alService *AlogService) GetAlog(id uint) (al alog.Alog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&al).Error
	return
}

// GetAlogInfoList 分页获取金额纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (alService *AlogService) GetAlogInfoList(info alogReq.AlogSearch) (list []*alog.Alog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alog.Alog{})
	var als []*alog.Alog
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.User_id != nil {
		db = db.Where("user_id = ?", info.User_id)
	}
	if info.Amount_type != nil {
		db = db.Where("amount_type = ?", info.Amount_type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&als).Error

	for _, al := range als {
		var u users.Users
		global.GVA_DB.Where("id = ?", *al.User_id).First(&u)
		al.User = u
	}
	return als, total, err
}
