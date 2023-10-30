import service from '@/utils/request'

// @Tags Config
// @Summary 创建配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "创建配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /config/createConfig [post]
export const createConfig = (data) => {
  return service({
    url: '/config/createConfig',
    method: 'post',
    data
  })
}

// @Tags Config
// @Summary 删除配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "删除配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /config/deleteConfig [delete]
export const deleteConfig = (data) => {
  return service({
    url: '/config/deleteConfig',
    method: 'delete',
    data
  })
}

// @Tags Config
// @Summary 批量删除配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /config/deleteConfig [delete]
export const deleteConfigByIds = (data) => {
  return service({
    url: '/config/deleteConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags Config
// @Summary 更新配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "更新配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /config/updateConfig [put]
export const updateConfig = (data) => {
  return service({
    url: '/config/updateConfig',
    method: 'put',
    data
  })
}

// @Tags Config
// @Summary 用id查询配置管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Config true "用id查询配置管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /config/findConfig [get]
export const findConfig = (params) => {
  return service({
    url: '/config/findConfig',
    method: 'get',
    params
  })
}

// @Tags Config
// @Summary 分页获取配置管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取配置管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /config/getConfigList [get]
export const getConfigList = (params) => {
  return service({
    url: '/config/getConfigList',
    method: 'get',
    params
  })
}
