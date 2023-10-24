import service from '@/utils/request'

// @Tags ArticleCategory
// @Summary 创建文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleCategory true "创建文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ac/createArticleCategory [post]
export const createArticleCategory = (data) => {
  return service({
    url: '/ac/createArticleCategory',
    method: 'post',
    data
  })
}

// @Tags ArticleCategory
// @Summary 删除文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleCategory true "删除文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteArticleCategory [delete]
export const deleteArticleCategory = (data) => {
  return service({
    url: '/ac/deleteArticleCategory',
    method: 'delete',
    data
  })
}

// @Tags ArticleCategory
// @Summary 批量删除文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ac/deleteArticleCategory [delete]
export const deleteArticleCategoryByIds = (data) => {
  return service({
    url: '/ac/deleteArticleCategoryByIds',
    method: 'delete',
    data
  })
}

// @Tags ArticleCategory
// @Summary 更新文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleCategory true "更新文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ac/updateArticleCategory [put]
export const updateArticleCategory = (data) => {
  return service({
    url: '/ac/updateArticleCategory',
    method: 'put',
    data
  })
}

// @Tags ArticleCategory
// @Summary 用id查询文章分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ArticleCategory true "用id查询文章分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ac/findArticleCategory [get]
export const findArticleCategory = (params) => {
  return service({
    url: '/ac/findArticleCategory',
    method: 'get',
    params
  })
}

// @Tags ArticleCategory
// @Summary 分页获取文章分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文章分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ac/getArticleCategoryList [get]
export const getArticleCategoryList = (params) => {
  return service({
    url: '/ac/getArticleCategoryList',
    method: 'get',
    params
  })
}
