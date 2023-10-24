package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

// InitArticleRouter 初始化 文章 路由信息
func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	aRouter := Router.Group("a").Use(middleware.OperationRecord())
	aRouterWithoutRecord := Router.Group("a")
	var aApi = v1.ApiGroupApp.ArticleApiGroup.ArticleApi
	{
		aRouter.POST("createArticle", aApi.CreateArticle)   // 新建文章
		aRouter.DELETE("deleteArticle", aApi.DeleteArticle) // 删除文章
		aRouter.DELETE("deleteArticleByIds", aApi.DeleteArticleByIds) // 批量删除文章
		aRouter.PUT("updateArticle", aApi.UpdateArticle)    // 更新文章
	}
	{
		aRouterWithoutRecord.GET("findArticle", aApi.FindArticle)        // 根据ID获取文章
		aRouterWithoutRecord.GET("getArticleList", aApi.GetArticleList)  // 获取文章列表
	}
}
