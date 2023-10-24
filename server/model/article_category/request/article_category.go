package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/article_category"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleCategorySearch struct{
    article_category.ArticleCategory
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
