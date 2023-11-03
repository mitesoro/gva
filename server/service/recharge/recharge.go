package recharge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/recharge"
	rechargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/recharge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/users"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type RechargeService struct {
}

// CreateRecharge 创建充值记录
// Author [piexlmax](https://github.com/piexlmax)
func (rgService *RechargeService) CreateRecharge(rg *recharge.Recharge) (err error) {
	a := *rg.Amount * 100
	rg.Amount = &a
	err = global.GVA_DB.Create(rg).Error
	if err == nil {
		var u users.Users
		if err = global.GVA_DB.Where("id = ?", rg.User_id).First(&u).Error; err == nil {
			u.Amount += *rg.Amount
			u.AvailableAmount += *rg.Amount
			if err = global.GVA_DB.Save(&u).Error; err != nil {
				global.GVA_LOG.Error("CreateRecharge.SaveUser err", zap.Error(err))
			}
			utils.AddAmountLog(int(u.ID), *rg.Amount, u.Amount, 1)
		}
	}
	return err
}

// DeleteRecharge 删除充值记录
// Author [piexlmax](https://github.com/piexlmax)
func (rgService *RechargeService) DeleteRecharge(rg recharge.Recharge) (err error) {
	err = global.GVA_DB.Delete(&rg).Error
	return err
}

// DeleteRechargeByIds 批量删除充值记录
// Author [piexlmax](https://github.com/piexlmax)
func (rgService *RechargeService) DeleteRechargeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]recharge.Recharge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRecharge 更新充值记录
// Author [piexlmax](https://github.com/piexlmax)
func (rgService *RechargeService) UpdateRecharge(rg recharge.Recharge) (err error) {
	err = global.GVA_DB.Save(&rg).Error
	return err
}

// GetRecharge 根据id获取充值记录
// Author [piexlmax](https://github.com/piexlmax)
func (rgService *RechargeService) GetRecharge(id uint) (rg recharge.Recharge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rg).Error
	return
}

// GetRechargeInfoList 分页获取充值记录
// Author [piexlmax](https://github.com/piexlmax)
func (rgService *RechargeService) GetRechargeInfoList(info rechargeReq.RechargeSearch) (list []recharge.Recharge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&recharge.Recharge{})
	var rgs []recharge.Recharge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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

	err = db.Order("id DESC").Find(&rgs).Error
	return rgs, total, err
}
