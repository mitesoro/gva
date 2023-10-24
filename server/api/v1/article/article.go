package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/article"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    articleReq "github.com/flipped-aurora/gin-vue-admin/server/model/article/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ArticleApi struct {
}

var aService = service.ServiceGroupApp.ArticleServiceGroup.ArticleService


// CreateArticle 创建文章
// @Tags Article
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article.Article true "创建文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /a/createArticle [post]
func (aApi *ArticleApi) CreateArticle(c *gin.Context) {
	var a article.Article
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Title":{utils.NotEmpty()},
        "Content":{utils.NotEmpty()},
        "Article_category":{utils.NotEmpty()},
    }
	if err := utils.Verify(a, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := aService.CreateArticle(&a); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteArticle 删除文章
// @Tags Article
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article.Article true "删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /a/deleteArticle [delete]
func (aApi *ArticleApi) DeleteArticle(c *gin.Context) {
	var a article.Article
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := aService.DeleteArticle(a); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteArticleByIds 批量删除文章
// @Tags Article
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /a/deleteArticleByIds [delete]
func (aApi *ArticleApi) DeleteArticleByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := aService.DeleteArticleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateArticle 更新文章
// @Tags Article
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body article.Article true "更新文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /a/updateArticle [put]
func (aApi *ArticleApi) UpdateArticle(c *gin.Context) {
	var a article.Article
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Title":{utils.NotEmpty()},
          "Content":{utils.NotEmpty()},
          "Article_category":{utils.NotEmpty()},
      }
    if err := utils.Verify(a, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := aService.UpdateArticle(a); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindArticle 用id查询文章
// @Tags Article
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query article.Article true "用id查询文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /a/findArticle [get]
func (aApi *ArticleApi) FindArticle(c *gin.Context) {
	var a article.Article
	err := c.ShouldBindQuery(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rea, err := aService.GetArticle(a.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rea": rea}, c)
	}
}

// GetArticleList 分页获取文章列表
// @Tags Article
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query articleReq.ArticleSearch true "分页获取文章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /a/getArticleList [get]
func (aApi *ArticleApi) GetArticleList(c *gin.Context) {
	var pageInfo articleReq.ArticleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := aService.GetArticleInfoList(pageInfo); err != nil {
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
