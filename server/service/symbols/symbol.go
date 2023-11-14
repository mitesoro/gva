package symbols

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/symbols"
	symbolsReq "github.com/flipped-aurora/gin-vue-admin/server/model/symbols/request"
)

type SymbolService struct {
}

// CreateSymbol 创建合约品种记录
// Author [piexlmax](https://github.com/piexlmax)
func (sbService *SymbolService) CreateSymbol(sb *symbols.Symbol) (err error) {
	sb.Amount = sb.Amount * 100
	err = global.GVA_DB.Create(sb).Error
	return err
}

// DeleteSymbol 删除合约品种记录
// Author [piexlmax](https://github.com/piexlmax)
func (sbService *SymbolService) DeleteSymbol(sb symbols.Symbol) (err error) {
	err = global.GVA_DB.Delete(&sb).Error
	return err
}

// DeleteSymbolByIds 批量删除合约品种记录
// Author [piexlmax](https://github.com/piexlmax)
func (sbService *SymbolService) DeleteSymbolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]symbols.Symbol{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSymbol 更新合约品种记录
// Author [piexlmax](https://github.com/piexlmax)
func (sbService *SymbolService) UpdateSymbol(sb symbols.Symbol) (err error) {
	sb.Amount = sb.Amount * 100
	err = global.GVA_DB.Save(&sb).Error
	return err
}

// GetSymbol 根据id获取合约品种记录
// Author [piexlmax](https://github.com/piexlmax)
func (sbService *SymbolService) GetSymbol(id uint) (sb symbols.Symbol, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sb).Error
	sb.Amount = sb.Amount / 100
	return
}

// GetSymbolInfoList 分页获取合约品种记录
// Author [piexlmax](https://github.com/piexlmax)
func (sbService *SymbolService) GetSymbolInfoList(info symbolsReq.SymbolSearch) (list []symbols.Symbol, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&symbols.Symbol{})
	var sbs []symbols.Symbol
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Code != "" {
		db = db.Where("code = ?", info.Code)
	}
	if info.Multiple != nil {
		db = db.Where("multiple = ?", info.Multiple)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&sbs).Error
	return sbs, total, err
}
