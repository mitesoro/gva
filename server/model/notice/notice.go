// 自动生成模板Notice
package notice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 公告 结构体  Notice
type Notice struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;"`  //标题 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`  //内容 
}


// TableName 公告 Notice自定义表名 notice
func (Notice) TableName() string {
  return "notice"
}

