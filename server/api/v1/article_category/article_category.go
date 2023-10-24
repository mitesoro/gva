package article_category

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/article_category"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    article_categoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/article_category/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ArticleCategoryApi struct {
}

var acService = service.ServiceGroupApp.Article_categoryServiceGroup.ArticleCategoryService


// CreateArticleCategory 创建文章分类
// @Tags ArticleCategory
// @Summary 创建文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article_category.ArticleCategory true "创建文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ac/createArticleCategory [post]
func (acApi *ArticleCategoryApi) CreateArticleCategory(c *gin.Context) {
	var ac article_category.ArticleCategory
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
    }
	if err := utils.Verify(ac, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := acService.CreateArticleCategory(&ac); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticleCategory 删除文章分类
// @Tags ArticleCategory
// @Summary 删除文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article_category.ArticleCategory true "删除文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteArticleCategory [delete]
func (acApi *ArticleCategoryApi) DeleteArticleCategory(c *gin.Context) {
	var ac article_category.ArticleCategory
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := acService.DeleteArticleCategory(ac); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleCategoryByIds 批量删除文章分类
// @Tags ArticleCategory
// @Summary 批量删除文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ac/deleteArticleCategoryByIds [delete]
func (acApi *ArticleCategoryApi) DeleteArticleCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := acService.DeleteArticleCategoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticleCategory 更新文章分类
// @Tags ArticleCategory
// @Summary 更新文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article_category.ArticleCategory true "更新文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateArticleCategory [put]
func (acApi *ArticleCategoryApi) UpdateArticleCategory(c *gin.Context) {
	var ac article_category.ArticleCategory
	err := c.ShouldBindJSON(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
      }
    if err := utils.Verify(ac, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := acService.UpdateArticleCategory(ac); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticleCategory 用id查询文章分类
// @Tags ArticleCategory
// @Summary 用id查询文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query article_category.ArticleCategory true "用id查询文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ac/findArticleCategory [get]
func (acApi *ArticleCategoryApi) FindArticleCategory(c *gin.Context) {
	var ac article_category.ArticleCategory
	err := c.ShouldBindQuery(&ac)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reac, err := acService.GetArticleCategory(ac.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reac": reac}, c)
	}
}

// GetArticleCategoryList 分页获取文章分类列表
// @Tags ArticleCategory
// @Summary 分页获取文章分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query article_categoryReq.ArticleCategorySearch true "分页获取文章分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/getArticleCategoryList [get]
func (acApi *ArticleCategoryApi) GetArticleCategoryList(c *gin.Context) {
	var pageInfo article_categoryReq.ArticleCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := acService.GetArticleCategoryInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
