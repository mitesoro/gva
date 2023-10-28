import service from '@/utils/request'

// @Tags Orders
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Orders true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /os/createOrders [post]
export const createOrders = (data) => {
  return service({
    url: '/os/createOrders',
    method: 'post',
    data
  })
}

// @Tags Orders
// @Summary 删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Orders true "删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /os/deleteOrders [delete]
export const deleteOrders = (data) => {
  return service({
    url: '/os/deleteOrders',
    method: 'delete',
    data
  })
}

// @Tags Orders
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /os/deleteOrders [delete]
export const deleteOrdersByIds = (data) => {
  return service({
    url: '/os/deleteOrdersByIds',
    method: 'delete',
    data
  })
}

// @Tags Orders
// @Summary 更新订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Orders true "更新订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /os/updateOrders [put]
export const updateOrders = (data) => {
  return service({
    url: '/os/updateOrders',
    method: 'put',
    data
  })
}

// @Tags Orders
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Orders true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /os/findOrders [get]
export const findOrders = (params) => {
  return service({
    url: '/os/findOrders',
    method: 'get',
    params
  })
}

// @Tags Orders
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /os/getOrdersList [get]
export const getOrdersList = (params) => {
  return service({
    url: '/os/getOrdersList',
    method: 'get',
    params
  })
}
