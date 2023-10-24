// 自动生成模板ArticleCategory
package article_category

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 文章分类 结构体  ArticleCategory
type ArticleCategory struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;"`  //名称 
}


// TableName 文章分类 ArticleCategory自定义表名 article_category
func (ArticleCategory) TableName() string {
  return "article_category"
}

