package article_category

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleCategoryRouter struct {
}

// InitArticleCategoryRouter 初始化 文章分类 路由信息
func (s *ArticleCategoryRouter) InitArticleCategoryRouter(Router *gin.RouterGroup) {
	acRouter := Router.Group("ac").Use(middleware.OperationRecord())
	acRouterWithoutRecord := Router.Group("ac")
	var acApi = v1.ApiGroupApp.Article_categoryApiGroup.ArticleCategoryApi
	{
		acRouter.POST("createArticleCategory", acApi.CreateArticleCategory)   // 新建文章分类
		acRouter.DELETE("deleteArticleCategory", acApi.DeleteArticleCategory) // 删除文章分类
		acRouter.DELETE("deleteArticleCategoryByIds", acApi.DeleteArticleCategoryByIds) // 批量删除文章分类
		acRouter.PUT("updateArticleCategory", acApi.UpdateArticleCategory)    // 更新文章分类
	}
	{
		acRouterWithoutRecord.GET("findArticleCategory", acApi.FindArticleCategory)        // 根据ID获取文章分类
		acRouterWithoutRecord.GET("getArticleCategoryList", acApi.GetArticleCategoryList)  // 获取文章分类列表
	}
}
