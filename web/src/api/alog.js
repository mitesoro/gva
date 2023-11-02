import service from '@/utils/request'

// @Tags Alog
// @Summary 创建金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alog true "创建金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /al/createAlog [post]
export const createAlog = (data) => {
  return service({
    url: '/al/createAlog',
    method: 'post',
    data
  })
}

// @Tags Alog
// @Summary 删除金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alog true "删除金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /al/deleteAlog [delete]
export const deleteAlog = (data) => {
  return service({
    url: '/al/deleteAlog',
    method: 'delete',
    data
  })
}

// @Tags Alog
// @Summary 批量删除金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /al/deleteAlog [delete]
export const deleteAlogByIds = (data) => {
  return service({
    url: '/al/deleteAlogByIds',
    method: 'delete',
    data
  })
}

// @Tags Alog
// @Summary 更新金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alog true "更新金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /al/updateAlog [put]
export const updateAlog = (data) => {
  return service({
    url: '/al/updateAlog',
    method: 'put',
    data
  })
}

// @Tags Alog
// @Summary 用id查询金额纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Alog true "用id查询金额纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /al/findAlog [get]
export const findAlog = (params) => {
  return service({
    url: '/al/findAlog',
    method: 'get',
    params
  })
}

// @Tags Alog
// @Summary 分页获取金额纪录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取金额纪录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /al/getAlogList [get]
export const getAlogList = (params) => {
  return service({
    url: '/al/getAlogList',
    method: 'get',
    params
  })
}
