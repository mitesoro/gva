import service from '@/utils/request'

// @Tags Symbol
// @Summary 创建合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Symbol true "创建合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sb/createSymbol [post]
export const createSymbol = (data) => {
  return service({
    url: '/sb/createSymbol',
    method: 'post',
    data
  })
}

// @Tags Symbol
// @Summary 删除合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Symbol true "删除合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sb/deleteSymbol [delete]
export const deleteSymbol = (data) => {
  return service({
    url: '/sb/deleteSymbol',
    method: 'delete',
    data
  })
}

// @Tags Symbol
// @Summary 批量删除合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sb/deleteSymbol [delete]
export const deleteSymbolByIds = (data) => {
  return service({
    url: '/sb/deleteSymbolByIds',
    method: 'delete',
    data
  })
}

// @Tags Symbol
// @Summary 更新合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Symbol true "更新合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sb/updateSymbol [put]
export const updateSymbol = (data) => {
  return service({
    url: '/sb/updateSymbol',
    method: 'put',
    data
  })
}

// @Tags Symbol
// @Summary 用id查询合约品种
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Symbol true "用id查询合约品种"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sb/findSymbol [get]
export const findSymbol = (params) => {
  return service({
    url: '/sb/findSymbol',
    method: 'get',
    params
  })
}

// @Tags Symbol
// @Summary 分页获取合约品种列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取合约品种列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sb/getSymbolList [get]
export const getSymbolList = (params) => {
  return service({
    url: '/sb/getSymbolList',
    method: 'get',
    params
  })
}
