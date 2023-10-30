// 自动生成模板Config
package configs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 配置管理 结构体  Config
type Config struct {
      global.GVA_MODEL
      Field  string `json:"field" form:"field" gorm:"column:field;comment:字段;"`  //字段 
      Value  string `json:"value" form:"value" gorm:"column:value;comment:值;"`  //值 
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;"`  //描述 
}


// TableName 配置管理 Config自定义表名 config
func (Config) TableName() string {
  return "config"
}

