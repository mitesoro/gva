// 自动生成模板Article
package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 文章 结构体  Article
type Article struct {
	global.GVA_MODEL
	Title            string `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;"`                           //标题
	Desc             string `json:"desc" form:"desc" gorm:"column:desc;comment:简介;size:255;"`                              //简介
	Content          string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`                    //内容
	Author           string `json:"author" form:"author" gorm:"column:author;comment:作者;"`                                 //作者
	Article_category *int   `json:"article_category" form:"article_category" gorm:"column:article_category;comment:文章分类;"` //文章分类
	Symbol           string `json:"symbol" form:"symbol" gorm:"column:symbol;comment:合约;size:255;"`                        // 合约
	IsRecommend      bool   `json:"is_recommend" form:"is_recommend" gorm:"column:is_recommend;comment:是否推荐;"`             //  是否推荐 1推荐0不推荐
}

// TableName 文章 Article自定义表名 article
func (Article) TableName() string {
	return "article"
}
