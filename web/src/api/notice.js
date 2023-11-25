import service from '@/utils/request'

// @Tags Notice
// @Summary 创建公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notice true "创建公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /no/createNotice [post]
export const createNotice = (data) => {
  return service({
    url: '/no/createNotice',
    method: 'post',
    data
  })
}

// @Tags Notice
// @Summary 删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notice true "删除公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /no/deleteNotice [delete]
export const deleteNotice = (data) => {
  return service({
    url: '/no/deleteNotice',
    method: 'delete',
    data
  })
}

// @Tags Notice
// @Summary 批量删除公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /no/deleteNotice [delete]
export const deleteNoticeByIds = (data) => {
  return service({
    url: '/no/deleteNoticeByIds',
    method: 'delete',
    data
  })
}

// @Tags Notice
// @Summary 更新公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Notice true "更新公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /no/updateNotice [put]
export const updateNotice = (data) => {
  return service({
    url: '/no/updateNotice',
    method: 'put',
    data
  })
}

// @Tags Notice
// @Summary 用id查询公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Notice true "用id查询公告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /no/findNotice [get]
export const findNotice = (params) => {
  return service({
    url: '/no/findNotice',
    method: 'get',
    params
  })
}

// @Tags Notice
// @Summary 分页获取公告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /no/getNoticeList [get]
export const getNoticeList = (params) => {
  return service({
    url: '/no/getNoticeList',
    method: 'get',
    params
  })
}
