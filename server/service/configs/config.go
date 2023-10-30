package configs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    configsReq "github.com/flipped-aurora/gin-vue-admin/server/model/configs/request"
)

type ConfigService struct {
}

// CreateConfig 创建配置管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (configService *ConfigService) CreateConfig(config *configs.Config) (err error) {
	err = global.GVA_DB.Create(config).Error
	return err
}

// DeleteConfig 删除配置管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (configService *ConfigService)DeleteConfig(config configs.Config) (err error) {
	err = global.GVA_DB.Delete(&config).Error
	return err
}

// DeleteConfigByIds 批量删除配置管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (configService *ConfigService)DeleteConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]configs.Config{},"id in ?",ids.Ids).Error
	return err
}

// UpdateConfig 更新配置管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (configService *ConfigService)UpdateConfig(config configs.Config) (err error) {
	err = global.GVA_DB.Save(&config).Error
	return err
}

// GetConfig 根据id获取配置管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (configService *ConfigService)GetConfig(id uint) (config configs.Config, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&config).Error
	return
}

// GetConfigInfoList 分页获取配置管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (configService *ConfigService)GetConfigInfoList(info configsReq.ConfigSearch) (list []configs.Config, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&configs.Config{})
    var configs []configs.Config
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Field != "" {
        db = db.Where("field = ?",info.Field)
    }
    if info.Value != "" {
        db = db.Where("value = ?",info.Value)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&configs).Error
	return  configs, total, err
}
