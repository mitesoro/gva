package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleSearch struct{
    article.Article
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
