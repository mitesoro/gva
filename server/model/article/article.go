// 自动生成模板Article
package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 文章 结构体  Article
type Article struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;"`  //标题 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`  //内容 
      Author  string `json:"author" form:"author" gorm:"column:author;comment:作者;"`  //作者 
      Article_category  *int `json:"article_category" form:"article_category" gorm:"column:article_category;comment:文章分类;"`  //文章分类 
}


// TableName 文章 Article自定义表名 article
func (Article) TableName() string {
  return "article"
}

