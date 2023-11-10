package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint             `gorm:"primarykey"` // 主键ID
	CreatedAt *model.LocalTime // 创建时间
	UpdatedAt *model.LocalTime // 更新时间
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"` // 删除时间
}
