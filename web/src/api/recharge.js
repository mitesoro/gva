import service from '@/utils/request'

// @Tags Recharge
// @Summary 创建充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recharge true "创建充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rg/createRecharge [post]
export const createRecharge = (data) => {
  return service({
    url: '/rg/createRecharge',
    method: 'post',
    data
  })
}

// @Tags Recharge
// @Summary 删除充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recharge true "删除充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rg/deleteRecharge [delete]
export const deleteRecharge = (data) => {
  return service({
    url: '/rg/deleteRecharge',
    method: 'delete',
    data
  })
}

// @Tags Recharge
// @Summary 批量删除充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rg/deleteRecharge [delete]
export const deleteRechargeByIds = (data) => {
  return service({
    url: '/rg/deleteRechargeByIds',
    method: 'delete',
    data
  })
}

// @Tags Recharge
// @Summary 更新充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Recharge true "更新充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rg/updateRecharge [put]
export const updateRecharge = (data) => {
  return service({
    url: '/rg/updateRecharge',
    method: 'put',
    data
  })
}

// @Tags Recharge
// @Summary 用id查询充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Recharge true "用id查询充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rg/findRecharge [get]
export const findRecharge = (params) => {
  return service({
    url: '/rg/findRecharge',
    method: 'get',
    params
  })
}

// @Tags Recharge
// @Summary 分页获取充值列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取充值列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rg/getRechargeList [get]
export const getRechargeList = (params) => {
  return service({
    url: '/rg/getRechargeList',
    method: 'get',
    params
  })
}
